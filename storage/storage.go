package storage

import (
	"github.com/tentsk8s/k8s-claimer/cluster"
	"github.com/tentsk8s/k8s-claimer/leases"
)

type Leaser interface {
	MarkCluster(leases.Lease)
	IsLeased(cluster.ID)
	ReleaseCluster(cluster.ID)
}

type ClusterFinder interface {
	FindAndMark(cluster.Type) error
}
type ClusterAdder interface {
	AddAndMark(cluster.Details)
}
