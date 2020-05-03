package leases

import (
	"fmt"

	"github.com/pborman/uuid"
)

type ID uuid.UUID

func (c *ID) String() string {
	return fmt.Sprintf("%s.%s", c.name, c.uid)
}

func (c *ID) MarshalJSON() ([]byte, error) {
	return []byte(c.String()), nil
}
