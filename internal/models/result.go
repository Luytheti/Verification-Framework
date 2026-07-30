package models

import "time"

type VerificationStatus string

const (
	StatusSuccess VerificationStatus = "SUCCESS"
	StatusFailed  VerificationStatus = "FAILED"
)

type VerificationResult struct {
	JobID     string
	Status    VerificationStatus
	Message   string
	Duration  time.Duration
	Completed time.Time
}
