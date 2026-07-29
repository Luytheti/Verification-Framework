# Components

This document describes every major component in PG Verify and its responsibilities.

---

# Discovery

## Responsibility

Discover PostgreSQL clusters available for verification.

## Inputs

None.

## Outputs

A collection of Cluster objects.

## Future Implementations

* Mock Discovery
* Kubernetes Discovery

---

# Scheduler

## Responsibility

Periodically initiate verification cycles.

## Responsibilities

* Trigger verification
* Create jobs
* Push jobs into the queue

The scheduler never performs verification itself.

---

# Job Queue

## Responsibility

Buffer verification jobs between producers and workers.

Responsibilities include:

* Receive jobs
* Hold pending jobs
* Provide jobs to workers

---

# Worker Pool

## Responsibility

Execute verification jobs concurrently.

Workers:

* Read jobs from the queue.
* Execute verification.
* Produce results.

Workers should never schedule new jobs.

---

# Verification Engine

## Responsibility

Run verification modules for a single cluster.

The engine coordinates execution but does not contain verification logic itself.

---

# Verifiers

Each verifier checks one aspect of recoverability.

Examples:

* Backup Verifier
* Storage Verifier
* Restore Verifier
* SQL Verifier
* WAL Verifier

All verifiers should implement a common interface.

---

# Result Aggregator

## Responsibility

Combine individual verifier results into a single verification report.

Responsibilities include:

* Overall status
* Duration
* Warnings
* Failures

---

# Storage

## Responsibility

Persist verification history.

Initial implementation:

* SQLite

Future implementation:

* PostgreSQL

---

# REST API

## Responsibility

Expose framework data to clients.

Responsibilities include:

* List clusters
* Retrieve verification history
* Retrieve latest status
* Trigger verification (future)

---

# Metrics

## Responsibility

Expose operational metrics.

Examples:

* Total jobs
* Failed jobs
* Queue depth
* Active workers
* Verification duration

---

# Logging

## Responsibility

Provide structured information about framework behaviour.

Logging should include:

* Startup
* Shutdown
* Scheduling
* Worker lifecycle
* Verification execution
* Errors

---

# Configuration

## Responsibility

Centralize framework configuration.

Examples:

* Worker count
* Verification interval
* Database path
* Log level
* Metrics port

Configuration should remain independent of business logic.
