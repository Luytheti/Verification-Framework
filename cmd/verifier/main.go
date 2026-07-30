package main

import (
	"fmt"

	"pg-verify/internal/discovery"
	"pg-verify/internal/job"
	"pg-verify/internal/models"
)

func main() {

	fmt.Println("===================================")
	fmt.Println("      PG Verify Framework")
	fmt.Println("===================================")

	fmt.Println()

	fmt.Println("Initializing discovery...")

	d := discovery.NewMockDiscovery()

	clusters, err := d.Discover()
	if err != nil {
		panic(err)
	}

	fmt.Printf("\nDiscovered %d PostgreSQL clusters\n\n", len(clusters))

	factory := job.NewFactory()

	jobs := make([]models.VerificationJob, 0)

	for _, cluster := range clusters {
		job := factory.Create(cluster)
		jobs = append(jobs, job)
	}
	fmt.Println()
	fmt.Printf("Created %d verification jobs\n\n", len(jobs))

	for _, job := range jobs {
		fmt.Printf(
			"Job %-25s Cluster: %-15s Namespace: %s\n",
			job.ID,
			job.Cluster.Name,
			job.Cluster.Namespace,
		)
	}
}
