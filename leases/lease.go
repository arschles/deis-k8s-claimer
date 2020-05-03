package leases

import (
	"encoding/json"
	"time"
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
	ClusterType    ClusterType `json:"cluster_type"`
	ClusterID      ClusterID   `json:"cluster_id"`
	ClusterName    string      `json:"cluster_name"`
	ExpirationTime *time.Time  `json:"lease_expiration_time"`
}

// NewLease creates a new lease with the given cluster name and expiration time
func NewLease(clusterName string, exprTime *time.Time) *Lease {
	return &Lease{
		ClusterName:    clusterName,
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

// ExpirationTime returns the expiration time of this lease.
func (l *Lease) IsExpired() bool {
	return l.ExpirationTime.After(time.Now())
}
