# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

**EH Hub Data Orchestration Platform** - A healthcare backend that orchestrates real-time patient health data flow from Garmin wearable devices through Garmin Connect Cloud to clinical monitoring systems (Vantiq) and patient compliance storage (Apple HealthKit).

- **Module Name**: `github.com/darkphotonKN/eh-hub-data-orchestration-platform`
- **Go Version**: 1.22.3
- **Framework**: Gin Web Framework
- **Purpose**: Healthcare data orchestration with dual-frequency processing
- **Compliance**: HIPAA-compliant patient data handling

## Architecture Summary

Single Go backend application serving as the central orchestration layer:

**Input**: Garmin Connect Cloud (Health API + webhooks)
**Output 1**: Vantiq Engine (real-time clinical monitoring)  
**Output 2**: iPhone EH Hub App → HealthKit (15min compliance batches)

**Key Data Types**: 5 critical health metrics for cardiac patients:
- SpO₂ (Blood Oxygen Saturation)
- ECG (Electrocardiogram) 
- Respiratory Rate
- Skin Temperature
- Heart Rate Variability (HRV)

## Development Commands

```bash
# Run with hot reload (requires Air)
make dev

# Build and run
make run

# Run all tests with coverage
make test

# Manual commands
go mod tidy                  # Update dependencies
go fmt ./...                 # Format code
```

## Project Structure

```
.
├── cmd/main.go              # Application entry point
├── config/                  # Configuration modules
│   └── routes.go           # Route configuration (webhook endpoints)
├── internal/               # Core application code
│   └── healthmetrics/      # Health data processing module
│       ├── handler.go      # Webhook handlers (Garmin ingestion)
│       ├── model.go        # Health data models and request/response types
│       └── service.go      # Business logic (validation, dual-frequency orchestration)
├── .env                    # Environment configuration
└── Makefile               # Build commands
```

## API Endpoints

### Webhook Endpoints
- `POST /api/webhooks/garmin` - Garmin Connect webhook for real-time health data

## Core Technical Constraints

1. **Stateless Processing**: No database - Vantiq handles persistence
2. **Dual Frequency Handling**: 
   - Real-time streaming to Vantiq for clinical monitoring
   - 15-minute batches to HealthKit for compliance
3. **HIPAA Compliance**: Patient ID masking, audit trails, secure processing
4. **High Availability**: Healthcare-grade reliability requirements

## Environment Configuration

```bash
PORT=6000
ENV=development
```

## Architecture Patterns

1. **Handler → Service Pattern**: Clean separation of HTTP concerns and business logic
2. **Stateless Processing**: No persistent storage in this orchestration layer
3. **Dual-Frequency Orchestration**: Same data processed at different frequencies for different consumers
4. **Healthcare Compliance**: Built-in patient ID masking and validation
5. **Error Resilience**: Graceful handling of downstream service failures

## Key Service Methods

### `ProcessGarminWebhook`
Core orchestration method that:
- Validates incoming Garmin health metrics
- Applies HIPAA-compliant patient ID masking
- Sends real-time data to Vantiq clinical monitoring
- Queues data for HealthKit 15-minute batching

### Data Flow
1. Garmin webhook → `HandleGarminWebhook` → `ProcessGarminWebhook`
2. Validation & transformation → HIPAA compliance
3. Parallel processing:
   - Real-time: `sendToVantiq()` for clinical alerts
   - Batched: `queueForHealthKit()` for compliance tracking

## Development Workflow

1. Start development server: `make dev`
2. Server runs on `http://localhost:6000`
3. Webhook endpoint: `http://localhost:6000/api/webhooks/garmin`
4. Register webhook URL with Garmin Connect Developer API

## Adding New Integrations

To add new health data sources (e.g., Apple Watch, Fitbit):

1. Create new webhook handler in `healthmetrics/handler.go`
2. Add source-specific models in `healthmetrics/model.go`
3. Extend validation logic in service for new data formats
4. Register new webhook routes in `config/routes.go`
5. Follow same dual-frequency orchestration pattern

## Testing Strategy

- Test webhook payload parsing and validation
- Mock external service calls (Vantiq, HealthKit)
- Validate HIPAA compliance (patient ID masking)
- Test error handling for downstream service failures
- Integration tests with sample Garmin webhook payloads

## Security & Compliance

- **Patient ID Masking**: Automatic anonymization of sensitive identifiers
- **Audit Logging**: All webhook processing logged for compliance
- **Input Validation**: Strict validation of health metric ranges
- **Error Handling**: Secure error responses without exposing patient data

## Notes

- No database persistence in orchestration layer (stateless design)
- Vantiq handles clinical data persistence and alerting
- HealthKit integration manages patient compliance tracking
- Built for healthcare-grade reliability and HIPAA compliance
- Focus on real-time processing with graceful degradation