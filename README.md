
# federated-multi-agent
**Federated Multi Agent Reinforcement Learning For Priority Aware Pod Scheduling In Kubernetes**

### Paper Information
- **Author(s):** Arunkumar Sambandam
- **Published In:** *********************************************International Journal of Leading Research Publication (IJLRP)
- **Publication Date:** ******Aug 2021
- **ISSN:** E-ISSN: **********2582-8010
- **DOI:**
- **Impact Factor:** *******9.56

### Abstract
This work addresses the limitations of Kubernetes’ centralized, heuristic-based scheduler in managing heterogeneous, priority-sensitive workloads at scale. It proposes a decentralized, 
federated multi-agent reinforcement learning framework that enables adaptive, priority-aware pod placement using local node intelligence. The approach optimizes scheduling decisions by 
jointly considering pod priority, resource efficiency, and scheduling latency while integrating seamlessly with existing Kubernetes APIs. Experimental evaluation on realistic microservices
workloads demonstrates improved priority satisfaction, throughput, and scalability compared to traditional and learning-based schedulers.

### Key Contributions
- **Federated Learning–Based Scheduler:**
  Proposed a decentralized reinforcement learning scheduler that replaces centralized, heuristic-driven Kubernetes scheduling.
  
- **Priority-Aware Placement:**
  Designed a learning objective that integrates pod priority, latency, and resource efficiency for improved placement decisions.
    
- **Decentralized and Coordinated Learning:**
  Implemented federated coordination to synchronize multiple local agents while preserving scalability and autonomy.
     
- **End-to-End Validation:**
  Built and evaluated a Kubernetes-native prototype showing consistent reductions in pod wait time across cluster sizes.
  
### Relevance & Real-World Impact
- **Reduced Scheduling Latency:**
  Significantly lowered pod waiting times, especially for high-priority workloads under contention.
   
- **Improved Scalability:**
Removed centralized bottlenecks, enabling faster and more adaptive scheduling as clusters grow.

- **Adaptive and Resilient Scheduling:**
    Continuously adjusted to dynamic workloads and node conditions, improving cluster stability.
  
  **Efficient Resource Usage:**
  Enhanced workload distribution and reduced congestion across heterogeneous nodes.
   
- **Production and Research Ready:**
    Delivered a deployable, API-compatible framework suitable for industry use, research, and advanced education.

### Experimental Results (Summary)

  | Nodes | Baseline scheduler laltenccy (ms) | Reinforced Adaptive wait Optimization (ms) | Improvment (%)  |
  |-------|-----------------------------------| -------------------------------------------| ----------------|
  | 3     |  4.8                              | 2.9                                        | 39.58           |
  | 5     |  4.3                              | 2.4                                        | 44.19           |
  | 7     |  3.9                              | 2                                          | 48.72           |
  | 9     |  3.6                              | 1.8                                        | 50.00           |
  | 11    |  3.4                              | 1.6                                        | 52.94           |

### Citation
Federated Multi Agent Reinforcement Learning For Priority Aware Pod Scheduling In Kubernetes
* Arunkumar Sambandam
* ***********************************International Journal of Leading Research Publication 
* ISSN E-ISSN: *****************************2582-8010
* License \
This research is shared for a academic and research purposes. For commercial use, please contact the author.\
**Resources** \
https://www.ijlrp.com*****************/ \
**Author Contact** \
**LinkedIn**: https://www.linkedin.com/**** | **Email**: arunkumar.sambandam@yahoo.com






