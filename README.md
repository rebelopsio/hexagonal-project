# Infrastructure Monitoring Dashboard Backend

A Go-based backend service implementing hexagonal architecture to aggregate infrastructure metrics from multiple sources and provide a unified API for monitoring dashboards.

## Project Overview

This project demonstrates hexagonal architecture (ports and adapters pattern) by creating a monitoring service that:
- Collects metrics from various infrastructure sources
- Aggregates and processes time-series data
- Provides multiple interfaces for accessing monitoring data
- Manages alerting and notification workflows

## Architecture Goals

### Hexagonal Architecture Implementation
- **Domain Layer**: Core business logic for metrics, alerts, and infrastructure monitoring
- **Primary Adapters**: REST API, gRPC service, CLI tool, webhook handlers
- **Secondary Adapters**: Database repositories, external service clients, caching layers
- **Ports**: Clean interfaces defining contracts between layers

### Key Learning Objectives
- Implement dependency inversion with Go interfaces
- Separate business logic from infrastructure concerns
- Create testable code with mockable dependencies
- Handle multiple data sources through unified abstractions
- Practice clean layering and component isolation

## Domain Model

### Core Entities
- **Metric**: Time-series data points with labels and metadata
- **Alert**: Configurable rules and thresholds for notifications
- **Resource**: Infrastructure components being monitored
- **HealthCheck**: Service availability and performance indicators

### Primary Use Cases
- Collect metrics from multiple sources
- Aggregate and transform time-series data
- Evaluate alert conditions and trigger notifications
- Generate monitoring reports and dashboards
- Manage metric retention and archival policies

## Adapter Implementations

### Primary Adapters (Inbound)
- **REST API**: HTTP endpoints for dashboard queries and configuration
- **gRPC Service**: Real-time metric streaming for live dashboards
- **CLI Tool**: Administrative interface for service management
- **Webhook Handler**: Endpoints for receiving external alerts

### Secondary Adapters (Outbound)
- **Prometheus Client**: Query metrics from Prometheus instances
- **CloudWatch Client**: Fetch AWS infrastructure metrics
- **Kubernetes Client**: Collect cluster and pod metrics
- **PostgreSQL Repository**: Store configuration and metadata
- **ClickHouse Repository**: Time-series data warehouse
- **Redis Cache**: Fast access to frequently queried metrics
- **SMTP/Slack Notifiers**: Alert delivery mechanisms

## Technical Requirements

### Core Dependencies
- Go 1.21+
- PostgreSQL 14+ (metadata storage)
- ClickHouse (time-series analytics)
- Redis (caching layer)

### External Integrations
- Prometheus/Grafana compatible APIs
- AWS CloudWatch SDK
- Kubernetes client-go
- SMTP/Slack webhooks for notifications

## Implementation Phases

### Phase 1: Domain and Ports
- [ ] Define core domain entities and value objects
- [ ] Create port interfaces for all external dependencies
- [ ] Implement core business logic without external dependencies
- [ ] Write comprehensive unit tests for domain layer

### Phase 2: Basic Adapters
- [ ] Implement in-memory repositories for testing
- [ ] Create basic REST API adapter
- [ ] Add PostgreSQL repository adapter
- [ ] Build simple CLI interface

### Phase 3: Data Sources
- [ ] Implement Prometheus metrics adapter
- [ ] Add CloudWatch integration
- [ ] Create Kubernetes metrics collector
- [ ] Add ClickHouse time-series adapter

### Phase 4: Advanced Features
- [ ] Real-time gRPC streaming service
- [ ] Redis caching layer
- [ ] Alert evaluation engine
- [ ] Notification delivery system

### Phase 5: Production Readiness
- [ ] Comprehensive error handling
- [ ] Observability and logging
- [ ] Configuration management
- [ ] Docker containerization
- [ ] Kubernetes deployment manifests

## Project Structure

```
infrastructure-monitoring/
├── cmd/
│   ├── api/           # REST API server
│   ├── grpc/          # gRPC streaming server
│   └── cli/           # Administrative CLI tool
├── internal/
│   ├── domain/        # Core business logic and entities
│   │   ├── metric/
│   │   ├── alert/
│   │   └── resource/
│   ├── ports/         # Interface definitions
│   │   ├── repositories/
│   │   ├── services/
│   │   └── notifiers/
│   └── adapters/      # External integrations
│       ├── api/       # REST/gRPC handlers
│       ├── cli/       # Command line interface
│       ├── repositories/
│       ├── clients/   # External service clients
│       └── notifiers/
├── pkg/
│   └── common/        # Shared utilities
├── configs/           # Configuration files
├── deployments/       # Kubernetes manifests
├── docs/              # Architecture documentation
└── tests/
    ├── integration/
    └── e2e/
```

## Success Criteria

### Architecture Quality
- Clean separation between domain and infrastructure layers
- All external dependencies accessed through interfaces
- No direct imports of infrastructure packages in domain code
- Comprehensive unit tests without external dependencies

### Functional Requirements
- Support multiple metric sources simultaneously
- Handle time-series data aggregation and querying
- Implement configurable alerting with multiple notification channels
- Provide both synchronous API and real-time streaming interfaces

### Technical Excellence
- Graceful error handling and recovery
- Comprehensive logging and metrics
- Proper configuration management
- Production-ready deployment setup

## Getting Started

1. **Study the Domain**: Understand monitoring concepts and time-series data patterns
2. **Design the Ports**: Define clean interfaces before implementing adapters
3. **Test-First Development**: Write tests for domain logic before implementation
4. **Incremental Building**: Start with in-memory adapters, then add real integrations
5. **Measure Success**: Evaluate how easily you can swap implementations

## Learning Resources

- Hexagonal Architecture principles and patterns
- Go interface design best practices
- Time-series database concepts
- Infrastructure monitoring fundamentals
- Kubernetes metrics and observability

---

**Note**: This is a learning project focused on architectural patterns. Start simple, iterate frequently, and prioritize clean separation of concerns over feature completeness.
