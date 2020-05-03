package gke

import (
	container "google.golang.org/api/container/v1"
)

// Lister implements the providers.Lister interface.
// it lists all the GKE clusters we have ownership over.
//
// This interface comes with a concrete implementation
// and a mock. Use it throughout your code so you can
// more easily unit test
type Lister interface {
	// List lists all of the clusters in the given project and zone
	List() (*container.ListClustersResponse, error)
}

// GKEClusterLister is a ClusterLister implementation that uses the GKE Go SDK to list clusters
// on a live GKE cluster
type gkeLister struct {
	svc    *container.Service
	projID string
	zone   string
}

// NewLister creates a new lister
func newLister(
	svc *container.Service,
	projID,
	zone string,
) *gkeLister {
	return &gkeLister{
		svc:    svc,
		projID: projID,
		zone:   zone,
	}
}

// List is the ClusterLister interface implementation
func (g *gkeLister) List() (*container.ListClustersResponse, error) {
	return g.svc.Projects.Zones.Clusters.List(
		g.projID,
		g.zone,
	).Do()
}
