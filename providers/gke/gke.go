package gke

// import (
// 	"encoding/json"
// 	"log"
// 	"net/http"
// 	"time"

// 	// "k8s.io/client-go/pkg/api/v1"

// 	"github.com/pborman/uuid"
// 	"github.com/tentsk8s/k8s-claimer/htp"
// 	"github.com/tentsk8s/k8s-claimer/leases"
// )

// // Lease will search for an available cluster on GKE which
// // matches the parameters passed in on the request
// // It will write back on the response the necessary connection
// // information in json format
// func Lease(
// 	projID,
// 	zone string,
// ) {
// 	clusterDetails, err := leaseOrCreate(&clusterDetails{
// 		projID: projID,
// 		zone:   zone,
// 	})

// 	availableCluster, err := searchForFreeCluster(clusterMap, leaseMap, req.ClusterRegex, req.ClusterVersion)
// 	if err != nil {
// 		switch e := err.(type) {
// 		case errNoAvailableOrExpiredClustersFound:
// 			log.Printf("No available clusters found")
// 			htp.Error(w, http.StatusConflict, "No available clusters found")
// 			return
// 		case errExpiredLeaseGKEMissing:
// 			log.Printf("Cluster %s has an expired lease but doesn't exist in GKE", e.clusterName)
// 			htp.Error(w, http.StatusInternalServerError, "Cluster %s has an expired lease but doesn't exist in GKE", e.clusterName)
// 			return
// 		default:
// 			log.Printf("Unknown error %s", e.Error())
// 			htp.Error(w, http.StatusInternalServerError, "Unknown error %s", e.Error())
// 			return
// 		}
// 	}

// 	newToken := uuid.NewUUID()
// 	kubeConfig, err := k8s.CreateKubeConfigFromCluster(availableCluster)
// 	if err != nil {
// 		log.Printf("Error creating kubeconfig file for cluster %s -- %s", availableCluster.Name, err)
// 		htp.Error(w, http.StatusInternalServerError, "Error creating kubeconfig file for cluster %s -- %s", availableCluster.Name, err)
// 		return
// 	}

// 	kubeConfigStr, err := k8s.MarshalAndEncodeKubeConfig(kubeConfig)
// 	if err != nil {
// 		log.Printf("Error marshaling & encoding kubeconfig -- %s", err)
// 		htp.Error(w, http.StatusInternalServerError, "Error marshaling & encoding kubeconfig -- %s", err)
// 		return
// 	}

// 	resp := api.CreateLeaseResp{
// 		KubeConfigStr:  kubeConfigStr,
// 		IP:             availableCluster.Endpoint,
// 		Token:          newToken.String(),
// 		ClusterName:    availableCluster.Name,
// 		ClusterVersion: availableCluster.CurrentNodeVersion,
// 	}

// 	leaseMap.CreateLease(newToken, leases.NewLease(availableCluster.Name, req.ExpirationTime(time.Now())))
// 	if err := k8s.SaveAnnotations(services, svc, leaseMap); err != nil {
// 		log.Printf("Error saving new lease to Kubernetes annotations -- %s", err)
// 		htp.Error(w, http.StatusInternalServerError, "Error saving new lease to Kubernetes annotations -- %s", err)
// 		return
// 	}

// 	if err := json.NewEncoder(w).Encode(resp); err != nil {
// 		log.Printf("Error encoding json -- %s", err)
// 		htp.Error(w, http.StatusInternalServerError, "Error encoding json -- %s", err)
// 		return
// 	}
// }
