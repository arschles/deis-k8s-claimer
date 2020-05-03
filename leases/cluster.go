package leases

import (
	"fmt"

	"github.com/pborman/uuid"
)

type ClusterType string

const (
	AzureCluster ClusterType = "azure"
	GKECluster   ClusterType = "gke"
)

func (c *ClusterType) UnmarshalJSON(b []byte) error {
	bStr := string(b)
	switch bStr {
	case string(AzureCluster):
		*c = AzureCluster
		return nil
	case string(GKECluster):
		*c = GKECluster
		return nil
	default:
		return fmt.Errorf("Unsupported cluster type: %s", bStr)
	}
}

type ClusterID struct {
	uid  uuid.UUID
	name string
}

func (c *ClusterID) String() string {
	return fmt.Sprintf("%s.%s", c.name, c.uid)
}

func (c *ClusterID) MarshalJSON() ([]byte, error) {
	return []byte(c.String()), nil
}
