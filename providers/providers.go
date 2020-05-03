import "github.com/tentsk8s/k8s-claimer/leases"
type Lister interface {
	ListAll() []*Details	
}

type Details interface {
	ClusterID leases.ClusterID
	ClusterType leases.ClusterType
}
