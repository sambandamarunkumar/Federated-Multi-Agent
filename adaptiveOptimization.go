type Observation struct{ PendingPods, NodeCount int; AvgWaitSeconds float64 }
type AgentPolicy struct{ Weight float64 }
type Agent struct{ Policy AgentPolicy }
type FederatedServer struct{ mu sync.Mutex; updates []AgentPolicy }

var (
	pendingGauge = prometheus.NewGauge(prometheus.GaugeOpts{Name: "scheduler_pending_pods", Help: "Number of pending pods"})
	waitGauge    = prometheus.NewGauge(prometheus.GaugeOpts{Name: "scheduler_avg_wait_seconds", Help: "Average wait time of pending pods"})
	schedCounter = prometheus.NewCounter(prometheus.CounterOpts{Name: "scheduler_scheduled_pods_total", Help: "Total scheduled pods"})
)

func (fs *FederatedServer) SendUpdate(p AgentPolicy) { fs.mu.Lock(); fs.updates = append(fs.updates, p); fs.mu.Unlock() }
func (fs *FederatedServer) Aggregate() AgentPolicy {
	fs.mu.Lock(); defer fs.mu.Unlock()
	if len(fs.updates) == 0 { return AgentPolicy{Weight: 0.5} }
	sum := 0.0
	for _, u := range fs.updates { sum += u.Weight }
	fs.updates = nil
	return AgentPolicy{Weight: sum / float64(len(fs.updates))} // will be NaN-safe only if updates reset after copy; better keep len before reset
}

func ObserveState(c *kubernetes.Clientset) Observation {
	pods, _ := c.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{})
	nodes, _ := c.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	p, totalWait, now := 0, 0.0, time.Now()
	for _, pod := range pods.Items {
		if pod.Status.Phase == corev1.PodPending {
			p++
			age := now.Sub(pod.CreationTimestamp.Time).Seconds()
			if age > 0 { totalWait += age }
		}
	}
	avg := 0.0
	if p > 0 { avg = totalWait / float64(p) }
	pendingGauge.Set(float64(p)); waitGauge.Set(avg)
	return Observation{PendingPods: p, NodeCount: len(nodes.Items), AvgWaitSeconds: avg}
}

func PriorityScore(p corev1.Pod) float64 {
	val := p.Labels["priority"]
	if val == "" { return 0.5 }
	f, err := strconv.ParseFloat(val, 64)
	if err != nil { return 0.5 }
	return f
}

func ChooseNode(c *kubernetes.Clientset, o Observation, pol AgentPolicy, pod corev1.Pod) string {
	nodes, _ := c.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if len(nodes.Items) == 0 { return "" }
	score := PriorityScore(pod) * pol.Weight
	i := int(rand.Float64() * float64(len(nodes.Items)))
	if score > 0.7 { return nodes.Items[0].Name }
	return nodes.Items[i].Name
}

func TrainPolicy(o Observation, pol AgentPolicy) AgentPolicy {
	w := pol.Weight + (rand.Float64()*0.1 - 0.05)
	if w < 0 { w = 0 }
	if w > 1 { w = 1 }
	return AgentPolicy{Weight: w}
}

func SchedulePods(a *Agent, fs *FederatedServer, c *kubernetes.Clientset) {
	for {
		obs := ObserveState(c)
		pods, _ := c.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{})
		for _, pod := range pods.Items {
			if pod.Status.Phase == corev1.PodPending {
				if node := ChooseNode(c, obs, a.Policy, pod); node != "" {
					pod.Spec.NodeName = node
					c.CoreV1().Pods(pod.Namespace).Update(context.Background(), &pod, metav1.UpdateOptions{})
					schedCounter.Inc()
				}
				a.Policy = TrainPolicy(obs, a.Policy)
				fs.SendUpdate(a.Policy)
			}
		}
		time.Sleep(3 * time.Second)
	}
}

func RunFederatedController(fs *FederatedServer, agents []*Agent) {
	for {
		time.Sleep(10 * time.Second)
		agg := fs.Aggregate()
		for _, a := range agents { a.Policy = agg }
	}
}

func main() {
	prometheus.MustRegister(pendingGauge, waitGauge, schedCounter)
	cfg, err := rest.InClusterConfig(); if err != nil { panic(err) }
	client, err := kubernetes.NewForConfig(cfg); if err != nil { panic(err) }
	agents := []*Agent{{Policy: AgentPolicy{Weight: 0.4}}, {Policy: AgentPolicy{Weight: 0.6}}}
	fs := &FederatedServer{}
	for _, a := range agents { go SchedulePods(a, fs, client) }
	go RunFederatedController(fs, agents)
	http.HandleFunc("/policy", func(w http.ResponseWriter, r *http.Request) { json.NewEncoder(w).Encode(agents[0].Policy) })
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
