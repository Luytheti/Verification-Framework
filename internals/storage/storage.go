package storage

import "github.com/Luytheti/pg-verify/internal/models"

type Storage interface {
	Save(result models.VerificationResult) error
}
