package job

import (
	"fmt"
	"time"

	"pg-verify/internal/models"
)

type Factory struct{}

func NewFactory() *Factory {
	return &Factory{}
}

func (f *Factory) Create(cluster models.Cluster) models.VerificationJob {
	return models.VerificationJob{
		ID:        fmt.Sprintf("%s-%d", cluster.Name, time.Now().UnixNano()),
		Cluster:   cluster,
		CreatedAt: time.Now(),
	}
}
