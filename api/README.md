# Sharingan API

The Sharingan API is a RESTful backend service built with Go and Gin, providing endpoints for application monitoring and incident management.

## Getting Started

### Prerequisites

- Go 1.20+
- PostgreSQL 17
- Keycloak for authentication

### Development

```bash
# Run with live reloading using Air
air -c .air.toml

# Generate API documentation with Swag
swag init -g ./cmd/main.go -o api/docs
```

## API Documentation

API documentation is available via Swagger UI at `/swagger/index.html` when the API is running.

## Project Structure

```
api/
├── cmd/               # Main application entry point
├── api/docs/          # Auto-generated Swagger documentation
├── internal/          # Internal packages
│   ├── controllers/   # Request handlers
│   ├── core/          # Core application logic
│   ├── custom_logger/ # Custom logging implementation
│   ├── custom-errors/ # Error handling
│   ├── database/      # Database connection and operations
│   ├── helpers/       # Helper functions
│   ├── middlewares/   # HTTP middleware
│   ├── migrations/    # Database migrations
│   ├── models/        # Data models
│   └── routes/        # API route definitions
└── tests/             # Test files
```

## Key Dependencies

- [air](https://github.com/air-verse/air) : Used to live reload the application
- [gin-swagger](https://github.com/swaggo/gin-swagger) : Swagger documentation generation
    - [declarative comment docs](https://github.com/swaggo/swag/blob/master/README.md#how-to-use-it-with-gin)
- [Gocloack](https://github.com/Nerzal/gocloak) : Golang's keycloak client
- [Gin cors](https://github.com/gin-contrib/cors) : Gin middleware/handler to enable CORS support
- [Validator](https://github.com/go-playground/validator) : For struct validation
- [Zap logger](https://github.com/uber-go/zap) : To add custom logger to app
- [lumberjack](https://github.com/natefinch/lumberjack/tree/v2.2.1) : For writing logs to rolling files
