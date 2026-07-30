package models

import "time"

type VerificationJob struct {
	ID        string
	Cluster   Cluster
	CreatedAt time.Time
}
