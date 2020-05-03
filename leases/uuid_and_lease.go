package leases

import (
	"github.com/pborman/uuid"
)

// UUIDAndLease is a simple struct to encapsulate the token
// (which is a UUID) for a lease and the corresponding lease itself
type UUIDAndLease struct {
	UUID  uuid.UUID
	Lease *Lease
}

// NewUUIDAndLease creates a new UUIDAndLease struct populated with u
// and l for its UUID and Lease fields, respectively
func NewUUIDAndLease(u uuid.UUID, l *Lease) *UUIDAndLease {
	return &UUIDAndLease{UUID: u, Lease: l}
}
