# PG Verify

> A modular verification framework for PostgreSQL backups running on Kubernetes.

---

## Overview

PG Verify is a cloud-native verification framework that continuously evaluates the recoverability of PostgreSQL clusters.

Unlike traditional backup systems that only report whether a backup completed successfully, PG Verify focuses on answering a more important question:

> **"Can this PostgreSQL cluster actually be recovered if production fails right now?"**

The framework periodically discovers PostgreSQL clusters, schedules verification jobs, executes verification concurrently using Go worker pools, stores historical results, and exposes APIs and metrics for monitoring.

The project is heavily inspired by production database reliability systems while remaining simple enough to be built as a learning project.

---

# Why?

A successful backup does **not** guarantee a successful recovery.

Real-world failures include:

* Corrupted backup files
* Missing WAL archives
* Storage failures
* Restore failures
* Database startup failures
* Corrupted data after restore

PG Verify attempts to detect these problems before a disaster occurs.

---

# Goals

* Learn Go through a real systems project.
* Learn concurrent programming using worker pools.
* Learn cloud-native backend architecture.
* Learn Kubernetes APIs.
* Learn PostgreSQL backup verification concepts.
* Build a production-inspired, modular verification framework.

---

# Architecture

```
Discovery
     │
     ▼
Scheduler
     │
     ▼
Job Queue
     │
     ▼
Worker Pool
     │
     ▼
Verification Engine
     │
     ▼
Result Store
     │
     ├────────► REST API
     │
     └────────► Metrics
```

---

# Tech Stack

* Go
* Goroutines
* Channels
* Context
* SQLite
* Kubernetes (client-go)
* PostgreSQL
* Prometheus
* Docker

---

# Project Structure

```
pg-verify/

cmd/

internal/

pkg/

configs/

deployments/

docs/

README.md
```

---

# Learning Objectives

This project focuses on understanding:

* Go concurrency
* Interfaces
* Worker pools
* Producer-consumer systems
* Scheduler design
* Clean architecture
* PostgreSQL backups
* Kubernetes APIs
* Monitoring and observability
* Distributed systems fundamentals

---
