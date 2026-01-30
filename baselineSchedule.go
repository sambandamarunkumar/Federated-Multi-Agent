import (
	"context"
	"math/rand"
	"strconv"
	"time"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func priorityScore(pod v1.Pod) float64 {
	val := pod.Labels["priority"]
	if val == "" {
		return 0.5
	}
	f, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return 0.5
	}
	return f
}

func chooseNode(client *kubernetes.Clientset, pod v1.Pod) string {
	nodes, err := client.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil || len(nodes.Items) == 0 {
		return ""
	}
	score := priorityScore(pod)
	if score >= 0.8 {
		return nodes.Items[0].Name
	}
	if score >= 0.5 && len(nodes.Items) > 1 {
		return nodes.Items[1].Name
	}
	i := int(rand.Float64() * float64(len(nodes.Items)))
	return nodes.Items[i].Name
}

func schedulePendingPods(client *kubernetes.Clientset) {
	for {
		pods, err := client.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{})
		if err != nil {
			time.Sleep(3 * time.Second)
			continue
		}
		for _, pod := range pods.Items {
			if pod.Status.Phase == v1.PodPending && pod.Spec.NodeName == "" {
				node := chooseNode(client, pod)
				if node == "" {
					continue
				}
				pod.Spec.NodeName = node
				_, _ = client.CoreV1().Pods(pod.Namespace).Update(context.Background(), &pod, metav1.UpdateOptions{})
			}
		}
		time.Sleep(3 * time.Second)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	cfg, err := rest.InClusterConfig()
	if err != nil {
		panic(err)
	}
	client, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		panic(err)
	}
	schedulePendingPods(client)
}
