# SaveGen API

A Go application using the Gin framework.

## Prerequisites

- Go 1.21 or later

## Getting Started

1. Install dependencies:
```bash
go mod tidy
```

2. Run the application:
```bash
go run main.go
```

The server will start on port 8080. You can test the health check endpoint by visiting:
```
http://localhost:8080/health
```

## API Endpoints

- `GET /health` - Health check endpoint 