package azure

import (
	"log"

	"github.com/Azure/azure-sdk-for-go/arm/containerservice"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/tentsk8s/k8s-claimer/config"
)

type azureLister struct {
	config *config.Azure
}

func NewLister(cfg *config.Azure) providers.Lister {
	return &azureLister{
		config: cfg,
	}
}

// List is the ClusterLister interface implementation
func (a *azureLister) List() []providers.Details { //(*containerservice.ListResult, error) {
	bearerAuthorizer, err := NewBearerAuthorizer(
		a.Config,
		azure.PublicCloud.ResourceManagerEndpoint,
	)
	if err != nil {
		log.Printf("Error trying to create Bearer Authorizer: %s", err)
		return nil, err
	}

	csClient := containerservice.NewContainerServicesClient(a.Config.SubscriptionID)
	csClient.Authorizer = bearerAuthorizer
	listResult, err := csClient.List()
	if err != nil {
		log.Printf("Error trying to fetch Azure Cluster List: %s\n", err)
		return nil, err
	}
	return &listResult, nil
}
