package leases

import (
	"github.com/pborman/uuid"
)

type ID struct {
	uuid.UUID
}

func (c *ID) String() string {
	return c.UUID.String()
}

func (c *ID) MarshalJSON() ([]byte, error) {
	return []byte(c.String()), nil
}
