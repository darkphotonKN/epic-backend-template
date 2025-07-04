# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a clean, minimal Go-based REST API template using the Gin web framework with in-memory storage. It demonstrates a simple handler → service pattern without external dependencies like databases or authentication.

- **Module Name**: `github.com/darkphotonKN/eh-hub-data-orchestration-platform`
- **Go Version**: 1.22.3
- **Framework**: Gin Web Framework
- **Storage**: In-memory (no database)
- **Authentication**: None (removed for simplicity)

## Development Commands

```bash
# Run with hot reload (requires Air)
make dev

# Build and run
make run

# Run all tests with coverage
make test

# Run tests with HTML coverage preview
make test-preview

# Manual commands
go mod tidy                  # Update dependencies
go fmt ./...                 # Format code
```

## Project Structure

```
.
├── cmd/main.go           # Application entry point
├── config/               # Configuration modules
│   └── routes.go        # Route configuration
├── internal/            # Core application code
│   └── item/            # Item domain (example module)
│       ├── handler.go   # HTTP handlers
│       ├── model.go     # Data models
│       └── service.go   # Business logic (with in-memory storage)
├── .env                 # Environment configuration
└── Makefile            # Build commands
```

## API Routes

All routes are prefixed with `/api`:

### Item Routes (CRUD Example)
- `POST /items` - Create new item
- `GET /items` - Get all items
- `GET /items/:id` - Get item by ID
- `PUT /items/:id` - Update item by ID
- `DELETE /items/:id` - Delete item by ID

## Environment Configuration

Required `.env` variables:
```
PORT=6060
ENV=development
```

## Architecture Patterns

1. **Clean Architecture**: Clear separation of concerns with handlers and services
2. **Dependency Injection**: Services are injected into handlers via constructors
3. **In-Memory Storage**: No external dependencies, data stored in memory
4. **RESTful API**: Standard HTTP methods and status codes
5. **JSON Communication**: All requests and responses use JSON

## Development Workflow

1. Start development server: `make dev`
2. Server runs on `http://localhost:6060`
3. API available at `http://localhost:6060/api/items`

## Adding New Modules

To add a new domain (e.g., "product"):

1. Create `internal/product/` directory
2. Add `handler.go`, `service.go`, and `model.go` files
3. Follow the same pattern as the `item` module:
   - Handler: HTTP request/response handling
   - Service: Business logic and in-memory storage
   - Model: Data structures and request/response types
4. Register routes in `config/routes.go`

## Testing Strategy

- Unit tests should be colocated with code files
- Test handlers and services separately
- Run `make test` to execute all tests with coverage
- Use table-driven tests for comprehensive coverage

## Notes

- No database required - all data stored in memory
- Data persists only during application runtime
- Air is used for hot reload during development
- No authentication or authorization (kept simple)
- UUID is used for entity IDs
- Follow Go naming conventions and patterns