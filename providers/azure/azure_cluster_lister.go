package azure

import (
	"github.com/tentsk8s/k8s-claimer/config"
)

// AzureClusterLister is a ClusterLister implementation that uses the Azure Go SDK to list clusters
// on a live Azure cluster
type AzureClusterLister struct {
	Config *config.Azure
}

// NewAzureClusterLister creates a new AzureClusterLister configured to use the given client.
func NewAzureClusterLister(azureConfig *config.Azure) *AzureClusterLister {
	return &AzureClusterLister{Config: azureConfig}
}
