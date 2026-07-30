package discovery

import "pg-verify/internal/models"

type Discovery interface {
	Discover() ([]models.Cluster, error)
}
