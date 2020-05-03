package azure

import (
	"log"

	"github.com/Azure/azure-sdk-for-go/arm/containerservice"
	"github.com/Azure/go-autorest/autorest/azure"
)

type azureLister struct {
	listRes *containerservice.ListResult
}

func NewLister(listRes *containerservice.ListResult) providers.Lister {
	return &azureLister{
		listRes: listRes,
	}
}

// List is the ClusterLister interface implementation
func (a *azureLister) List() (*containerservice.ListResult, error) {
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
