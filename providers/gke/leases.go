package gke

import (
	"fmt"
	"time"

	"github.com/tentsk8s/k8s-claimer/cluster"
	"github.com/tentsk8s/k8s-claimer/leases"
	"github.com/tentsk8s/k8s-claimer/providers"
)

// NewLeaser creates a new leaser implementation
// for GKE
func NewLeaser(storage leaseStorage) providers.Leaser {
	return &leaseOps{storage: storage}
}

type leaseOps struct {
	storage leaseStorage
}

func (l *leaseOps) Acquire(
	cType cluster.Type,
	dur time.Duration,
) (*cluster.Details, error) {
	lease := l.storage.findAndLease(dur)
	if lease != nil {
		return leaseToStdDetails(lease), nil
	}
	// TODO: create a new cluster and call l.storage.addNewCluster
	return nil, nil
}

func (l *leaseOps) Free(id leases.ID) (*cluster.Details, error) {
	lease, err := l.storage.free(id)
	if err != nil {
		return nil, fmt.Errorf("Error releasing lease (%s)", err)
	}
	return leaseToStdDetails(lease), err
}

func leaseToStdDetails(lease *leases.Lease) *cluster.Details {
	return &cluster.Details{
		ID:   lease.ClusterDetails.ID,
		Type: cluster.GKECluster,
	}
}
