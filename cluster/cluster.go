package cluster

import (
	"fmt"

	"github.com/google/uuid"
)

type Type string

const (
	AzureCluster Type = "azure"
	GKECluster   Type = "gke"
)

func (c *Type) UnmarshalJSON(b []byte) error {
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

type Details struct {
	ID   ID
	Type Type
}

type ID uuid.UUID
