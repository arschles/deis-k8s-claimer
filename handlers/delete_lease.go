package handlers

import (
	"log"
	"net/http"

	"github.com/pborman/uuid"
	"github.com/tentsk8s/k8s-claimer/config"
	"github.com/tentsk8s/k8s-claimer/htp"
	"github.com/tentsk8s/k8s-claimer/providers/azure"
	"github.com/tentsk8s/k8s-claimer/providers/gke"
)

var (
	skipDeleteNamespaces = map[string]struct{}{
		"default":     struct{}{},
		"kube-system": struct{}{},
	}
)

// DeleteLease returns the http handler for the DELETE /lease/{token} endpoint
func DeleteLease(
	stg leases.Storage,
	clusterListers *
	gkeClusterLister gke.ClusterLister,
	azureClusterLister azure.ClusterLister,
	azureConfig *config.Azure,
	googleConfig *config.Google,
) http.Handler {
	return http.HandlerFunc(
		func(
			w http.ResponseWriter,
			r *http.Request,
		) {
			pathElts := htp.SplitPath(r)
			if len(pathElts) != 3 {
				log.Println("Path must be in the format /lease/{provider}/{token}")
				htp.Error(w, http.StatusBadRequest, "Path must be in the format /lease/{token}")
				return
			}

			provider := pathElts[1]
			leaseToken := uuid.Parse(pathElts[2])
			if leaseToken == nil {
				log.Printf("Lease token %s is invalid", pathElts[1])
				htp.Error(w, http.StatusBadRequest, "Lease token %s is invalid", pathElts[1])
				return
			}

			deletedErr := stg.DeleteCluster(lease)
			if deletedErr != nil {
				htp.Error(w, 300, "%s", deletedErr)
				return
			}
			w.WriteHeader(http.StatusOK)
		})
	}
}
