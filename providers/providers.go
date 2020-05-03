package providers

import (
	"time"

	"github.com/tentsk8s/k8s-claimer/cluster"
	"github.com/tentsk8s/k8s-claimer/leases"
)

type Leaser interface {
	Acquire(cluster.Type, time.Duration) (*cluster.Details, error)
	Free(leases.ID) (*cluster.Details, error)
}

/*
		switch req.CloudProvider {
		case "google":
			if googleConfig.ValidConfig() {
				gke.Lease(w, req, gkeClusterLister, services, k8sServiceName, googleConfig.ProjectID, googleConfig.Zone)
			} else {
				log.Println("Unable to satisfy this request because the Google provider is not properly configured.")
				htp.Error(w, http.StatusInternalServerError, "Unable to satisfy this request because the Google provider is not properly configured.")
			}
		case "azure":
			if azureConfig.ValidConfig() {
				azure.Lease(w, req, azureClusterLister, services, azureConfig, k8sServiceName)
			} else {
				log.Println("Unable to satisfy this request because the Azure provider is not properly configured.")
				htp.Error(w, http.StatusInternalServerError, "Unable to satisfy this request because the Azure provider is not properly configured.")
			}
		default:
			log.Printf("Unable to find suitable provider for this request -- Provider:%s", req.CloudProvider)
			htp.Error(w, http.StatusBadRequest, "Unable to find suitable provider for this request -- Provider:%s", req.CloudProvider)
		}
	})
*/
