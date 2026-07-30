package discovery

import "github.com/Luytheti/pg-verify/internal/models"

type Discovery interface {
	Discover() ([]models.Cluster, error)
}
