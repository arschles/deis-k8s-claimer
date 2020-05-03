package handlers

import (
	"net/http"

	"github.com/pborman/uuid"
	"github.com/tentsk8s/k8s-claimer/htp"
)

// DeleteLease is the handler for
// DELETE /leases/<id>
type DeleteLease struct {
	releaser providers.Releaser
}

func (d DeleteLease) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pathElts := htp.SplitPath(r)
	if len(pathElts) != 2 {
		htp.Error(
			w,
			http.StatusBadRequest,
			"You must specify the lease ID",
		)
		return
	}
	id := pathElts[1]
	uid := uuid.Parse(id)
	if uid == nil {
		htp.Error(
			w,
			http.StatusBadRequest,
			"The ID you passed is invalid",
		)
	}

	deletedErr := d.releaser.ReleaseLease(uid)
	if deletedErr != nil {
		htp.Error(
			w,
			http.StatusBadRequest,
			"Bad request (%s)",
			err,
		)
	}
	w.WriteHeader(200)
}

// DeleteLeaseReq is the encoding/json compatible struct that represents the DELETE /lease request body
type deleteLeaseReq struct {
	ID leases.ClusterID `json:"id"`
}
