package storage

import "pg-verify/internal/models"

type Storage interface {
	Save(result models.VerificationResult) error
}
