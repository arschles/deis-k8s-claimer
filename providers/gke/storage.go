package gke

import (
	"fmt"
	"sync"
	"time"

	"github.com/tentsk8s/k8s-claimer/cluster"
	"github.com/tentsk8s/k8s-claimer/leases"
)

type clusterDetails struct {
	stdDetails *cluster.Details
	projID     string
	zone       string
}

type leaseStorage struct {
	// map from lease ID to leases.
	// corresponds to clusters based on
	// lease.ClusterDetails.ID
	leases map[leases.ID]*leases.Lease
	// Lookup from cluster ID to cluster details
	clusters map[cluster.ID]*clusterDetails
	mut      *sync.Mutex
}

func (l leaseStorage) findAndLease(newDur time.Duration) *leases.Lease {
	l.mut.Lock()
	defer l.mut.Unlock()
	for id, lease := range l.leases {
		if lease.IsExpired() {
			lease.Renew(newDur)
			return lease
		}
	}
	return nil
}

func (l leaseStorage) free(id lease.ID) (*leases.Lease, error) {
	l.mut.Lock()
	defer l.mut.Unlock()
	lac, ok := l.leases[id]
	if !ok {
		return nil, fmt.Errorf("No lease exists with ID %s", id)
	}
	if lac.IsExpired() {
		return nil, fmt.Errorf("Lease %s was already expired", id)
	}
	lac.lease.Free()
	return lac, nil
}

func (l leaseStore) addCluster(cluster *clusterDetails) {
	l.mut.Lock()
	defer l.mut.Unlock()
	l.clusters[cluster.stdDetails.ID] = cluster
}

func (l leaseStore) addAndLease(
	cluster *clusterDetails
	lease *leases.Lease,
) {
	l.mut.Lock()
	defer l.mut.Unlock()
	l.addCluster(cluster)
	l.leases[lease.ID] = lease
}
