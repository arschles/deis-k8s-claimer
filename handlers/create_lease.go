package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/tentsk8s/k8s-claimer/htp"
	"github.com/tentsk8s/k8s-claimer/providers/azure"
	"github.com/tentsk8s/k8s-claimer/providers/gke"
)

// CreateLease creates the handler that responds to the POST
// /lease endpoint
type CreateLease struct {
	lister providers.Lister
	leaser providers.Leaser
}

func (c CreateLease) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	req := new(createLeaseReq)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("Error decoding JSON -- %s", err)
		htp.Error(
			w, http.StatusBadRequest,
			"Error decoding JSON (%s"),
			err,
		)
		return
	}

	details, err := c.leaser.Lease(req.ClusterType)
	if err != nil {
		htp.Error(
			w,
			http.StatusBadRequest,
			"Couldn't lease (%s)",
			err,
		)
		return
	}
	resp := &createLeaseResp {
		KubeConfigStr: "TODO",
		IP             : "TODO",
		ID details.ClusterID,
		ClusterType details.ClusterType,
		TimeLeft time.Duration(req.MaxTimeSec) * time.Seconds,
	}
	
	resBytes, err := json.Marshal(resp)
	if err != nil {
		htp.Error(
			w,
			http.StatusInternalServerError,
			"Couldn't return JSON (%s)",
			err,
		)
		return
	}
	w.Write(resBytes)
}


// createLeaseReq is the encoding/json compatible struct that
// represents the POST /lease request body
type createLeaseReq struct {
	MaxTimeSec  int                `json:"max_time"`
	ClusterType leases.ClusterType `json:"cloud_provider"`
}

// createLeaseResp is the encoding/json compatible struct that
// represents the POST /lease response body
type createLeaseResp struct {
	KubeConfigStr  string `json:"kubeconfig"`
	IP             string `json:"ip"`
	ID cluster.ID `json:"id"`
	ClusterType cluster.Type `json:"cluster_type"`
}

// KubeConfigBytes decodes c.KubeConfig by the RFC 4648 standard.
// See http://tools.ietf.org/html/rfc4648 for more information
func (c createLeaseResp) KubeConfigBytes() ([]byte, error) {
	kubeConfigBytes, err := base64.StdEncoding.DecodeString(c.KubeConfigStr)
	if err != nil {
		return nil, err
	}
	return kubeConfigBytes, nil
}

// KubeConfig returns decoded and unmarshalled Kubernetes client
// configuration
func (c createLeaseResp) KubeConfig() (*k8s.KubeConfig, error) {
	configBytes, err := c.KubeConfigBytes()
	if err != nil {
		return nil, err
	}
	kubeConfig := &k8s.KubeConfig{}
	if err := json.Unmarshal(configBytes, kubeConfig); err != nil {
		return nil, err
	}
	return kubeConfig, nil
}
