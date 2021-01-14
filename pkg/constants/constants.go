package constants

const BackTickHomeFile = "~/.kube/config"

// How to heuristically detect a kind / ekz cluster using labels
// "io.x-k8s.kind.cluster": "ekz"
// io.x-k8s.ekz.cluster=ekz
const (
	EKZClusterLabel  = "io.x-k8s.ekz.cluster"
	KINDClusterLabel = "io.x-k8s.kind.cluster"

	EKZNetworkLabel = "io.x-k8s.ekz.network"
)
