package leases

import (
	"encoding/json"
	"time"

	"github.com/tentsk8s/k8s-claimer/cluster"
)

const (
	// TimeFormat is the standard format that all lease expiration times are stored as in k8s
	// annotations
	TimeFormat = time.RFC3339
)

var (
	zeroTime time.Time
)

// Lease is the json-encodable struct that represents
// a lease for a cluster
type Lease struct {
	ID             ID
	ClusterDetails *cluster.Details `json:"cluster_details"`
	ExpirationTime *time.Time       `json:"lease_expiration_time"`
}

// NewLease creates a new lease with the given cluster name and expiration time
func NewLease(
	ID ID,
	clusterDetails *cluster.Details,
	exprTime *time.Time,
) *Lease {
	return &Lease{
		ID:             ID,
		ClusterDetails: clusterDetails,
		ExpirationTime: exprTime,
	}
}

// ParseLease decodes leaseStr from json into a Lease structure. Returns nil and any decoding error
// if there was one, and a valid lease and nil otherwise
func ParseLease(leaseStr string) (*Lease, error) {
	l := new(Lease)
	if err := json.Unmarshal([]byte(leaseStr), l); err != nil {
		return nil, err
	}
	return l, nil
}

// IsExpired returns true if the lease is expired
func (l *Lease) IsExpired() bool {
	return l.ExpirationTime.After(time.Now())
}

func (l *Lease) Free() error {
	l.ExpirationTime = &zeroTime
	return nil
}

func (l *Lease) Renew(newDur time.Duration) {
	newExpr := time.Now().Add(newDur)
	l.ExpirationTime = &newExpr
}
