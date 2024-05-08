
集群信息在pod和node中获取方式不同，ctrsploit通过实现两套信息收集方式实现集群信息收集。

其中 cluster.go 声明了 interface, node/node.go, pod/pod.go 分别是两种实现。