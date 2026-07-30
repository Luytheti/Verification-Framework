package verifier

import "pg-verify/internal/models"

type Verifier interface {
	Verify(job models.VerificationJob) (models.VerificationResult, error)
}
