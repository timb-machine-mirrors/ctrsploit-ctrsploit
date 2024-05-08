package cluster

type Cluster interface {
	GetApiServerAddr() (addr string, err error)
}
