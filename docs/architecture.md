# System Architecture

## High-Level Overview

PG Verify follows a modular, event-driven architecture composed of independent components.

Each component has a single responsibility and communicates through well-defined interfaces.

```
                  +----------------------+
                  |   PostgreSQL         |
                  |     Clusters         |
                  +----------+-----------+
                             |
                             |
                      Discovery Module
                             |
                             |
                      Scheduler Module
                             |
                             |
                         Job Queue
                             |
                   +---------+---------+
                   |                   |
             Worker Pool         Worker Pool
                   |                   |
                   +---------+---------+
                             |
                     Verification Engine
                             |
            +--------+--------+---------+
            |        |        |         |
        Backup    Storage   Restore   SQL
       Verifier   Verifier Verifier Verifier
                             |
                     Verification Result
                             |
                    Result Aggregator
                             |
                    SQLite / Repository
                             |
               +-------------+-------------+
               |                           |
           REST API                 Prometheus
```

---

# Architectural Principles

## Single Responsibility

Every module performs one clearly defined task.

Examples:

* Discovery only discovers clusters.
* Scheduler only schedules jobs.
* Workers only execute jobs.
* Verifiers only verify.

No component should have multiple unrelated responsibilities.

---

## Dependency Inversion

High-level components depend on interfaces rather than concrete implementations.

For example:

```
Scheduler

↓

Discovery Interface

↓

Mock Discovery

or

Kubernetes Discovery
```

Replacing the implementation should not require changes to the scheduler.

---

## Producer–Consumer Model

The scheduler acts as the producer.

Workers act as consumers.

```
Scheduler

↓

Job Queue

↓

Workers
```

This allows concurrency while limiting resource usage.

---

## Modular Verification

Verification consists of multiple independent checks.

```
Cluster

↓

Backup Check

↓

Storage Check

↓

Restore Check

↓

SQL Check

↓

Overall Result
```

Each verification module can be added, removed, or modified independently.

---

## Data Flow

Each verification cycle follows this sequence:

1. Scheduler triggers a verification cycle.
2. Discovery retrieves available PostgreSQL clusters.
3. A verification job is created for each cluster.
4. Jobs are placed into the queue.
5. Worker goroutines consume jobs.
6. Workers execute verification modules.
7. Results are aggregated.
8. Results are stored.
9. APIs and metrics expose the latest status.

---

# Why a Worker Pool?

Launching one goroutine per verification request is simple but can lead to excessive resource consumption.

Instead, the framework uses a bounded worker pool.

Advantages include:

* Predictable resource usage
* Better scalability
* Controlled concurrency
* Simpler monitoring
* Graceful shutdown

---

# Why Interfaces?

Interfaces make the framework extensible.

Examples include:

* Discovery implementations
* Storage implementations
* Verification modules

Future integrations should require new implementations rather than modifications to existing components.

---

# Future Architecture

Future versions will replace mocked components with production integrations.

Examples:

* Kubernetes Discovery
* CloudNativePG Backup CRDs
* Real Restore Verification
* Prometheus Metrics
* Grafana Dashboards
* OpenTelemetry Tracing

The overall architecture should remain unchanged as these implementations evolve.
