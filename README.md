
# federated-multi-agent
**FEDERATED MULTI AGENT REINFORCEMENT LEARNING FOR PRIORITY AWARE POD SCHEDULING IN KUBERNETES**

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
- **Improved Precision in Bottleneck Diagnosis:**
  Significantly reduced misattribution of performance issues by enabling precise identification of Input/Output bottlenecks through multimodal signal correlation.
   
- **Faster and More Reliable Performance Troubleshooting:**
Enabled quicker diagnosis and mitigation of performance degradation, reducing operational delays caused by fragmented observability and manual analysis.

- **Scalable Observability for Distributed and Cloud-Native Pipelines:**
    Demonstrated effectiveness across varying cluster sizes and workload intensities, maintaining diagnostic accuracy as coordination overhead and I/O contention evolve with scale.
  
  **Operational Stability Under Dynamic Workloads:**
  Improved system stability by providing continuous visibility into shifting I/O contention patterns, supporting proactive performance management instead of reactive tuning.
   
- **Production and Research Applicability:**
    Delivered a framework compatible with modern distributed and cloud native environments, offering a complete reference model architecture, implementation, and evaluation suitable
    for industry observability platforms, academic research, and advanced systems education.

### Experimental Results (Summary)

  | Nodes | Baseline Request Completion Time (ms) | Multimodal Request Completion Tim (ms) | Improvment (%)  |
  |-------|---------------------------------------| ---------------------------------------| ----------------|
  | 3     |  120                                  | 95                                     | 20.83           |
  | 5     |  145                                  | 115                                    | 20.69           |
  | 7     |  175                                  | 140                                    | 20.00           |
  | 9     |  210                                  | 165                                    | 21.43           |
  | 11    |  250                                  | 195                                    | 22.00           |

### Citation
Multimodal Observability for Input Output Bottleneck Detection
* Arunkumar Sambandam
* ***********************************International Journal of Leading Research Publication 
* ISSN E-ISSN: *****************************2582-8010
* License \
This research is shared for a academic and research purposes. For commercial use, please contact the author.\
**Resources** \
https://www.ijlrp.com*****************/ \
**Author Contact** \
**LinkedIn**: https://www.linkedin.com/**** | **Email**: arunkumar.sambandam@yahoo.com






