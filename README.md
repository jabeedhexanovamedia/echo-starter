# Echo Starter - Production Grade Go Boilerplate

[![Go](https://img.shields.io/badge/Go-1.24+-blue?logo=go&logoColor=white)](https://golang.org/)
[![Echo](https://img.shields.io/badge/Echo-v4-brightgreen)](https://echo.labstack.com/)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/jabeedhexanovamedia/echo-starter)](https://goreportcard.com/report/github.com/jabeedhexanovamedia/echo-starter)

A **production-grade Go + Echo boilerplate** for building scalable, testable, and maintainable backend applications.  
This starter template is **perfect for solo projects, team projects, or learning best practices in Go**.

---

## Features

- **Clean Architecture**: Well-organized folder structure following Go best practices
- **Thin Entrypoint**: Minimal `main.go` with clear separation of concerns
- **HTTP Server Lifecycle**: Isolated server logic from business logic
- **Layered Structure**: Handlers, services, repositories for clean code
- **Global Middleware**: Built-in logger and recovery middleware
- **Graceful Shutdown**: Ready for production deployments
- **Observability**: Logger, tracing, and metrics support
- **Background Jobs**: Folder for Asynq workers
- **Config Management**: Environment variables and `.env` support
- **Modular & Testable**: Easy to test and maintain

---

## 📁 Folder Structure

```
echo-starter/
├── cmd/
│   └── server/                  #  Entry point (main.go)
├── internal/
│   ├── app/                     #  Application wiring / Dependency Injection
│   ├── config/                  #  Configuration module
│   ├── errors/                  #  Custom typed errors
│   ├── handler/                 #  HTTP handlers
│   ├── jobs/                    #  Background workers
│   ├── middleware/              #  Custom middleware
│   ├── model/                   #  Domain models
│   ├── observability/           #  Logger / metrics / tracing
│   ├── repository/              #  Database access layer
│   ├── server/                  #  Server lifecycle & route registration
│   ├── service/                 #  Business logic
│   └── shutdown/                #  Graceful shutdown logic
├── migrations/                  #  Database migrations
├── configs/                     #  Configuration templates
├── scripts/                     #  Development / operations scripts
├── test/                        #  Integration / end-to-end tests
├── .env.example                 #  Environment variables template
├── Taskfile.yml                 #  Task automation
├── go.mod                       #  Go module file
├── go.sum                       #  Go dependencies checksum
└── README.md                    #  This file
```

---

## Quick Start

### Folder Creation & Initial Setup

If starting from scratch, create the project structure:

```bash
# Create the directory structure
mkdir -p cmd/server internal/{app,config,errors,handler,jobs,middleware,model,observability,repository,server,service,shutdown} migrations configs scripts test

# Initialize Go module
go mod init github.com/yourusername/echo-starter

# Install dependencies
go mod tidy
```

### Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/jabeedhexanovamedia/echo-starter.git
   cd echo-starter
   ```

2. **Install dependencies**:

   ```bash
   go mod download
   ```

3. **Set up environment** (optional):
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

### ▶ Running the Application

Start the server:

```bash
go run cmd/server/main.go
```

The server will start on `http://localhost:8080` by default.

For development with hot reload (if using Taskfile):

```bash
task dev
```

### Testing

Run tests:

```bash
go test ./...
```

For integration tests:

```bash
go test ./test/...
```

---

## Configuration

The application uses environment variables for configuration. Create a `.env` file in the root directory:

```env
# Server Configuration
PORT=8080
HOST=localhost

# Database
DB_HOST=localhost
DB_PORT=5432
DB_NAME=echo_starter
DB_USER=user
DB_PASSWORD=password

# Logging
LOG_LEVEL=info

# Other settings...
```

---

## Development

### Adding New Features

1. **Models**: Define your domain models in `internal/model/`
2. **Repositories**: Implement data access in `internal/repository/`
3. **Services**: Add business logic in `internal/service/`
4. **Handlers**: Create HTTP endpoints in `internal/handler/`
5. **Routes**: Register routes in `internal/server/server.go`

### Database Migrations

Add migration files to the `migrations/` directory and run them using your preferred migration tool.

### Background Jobs

Implement workers in `internal/jobs/` using Asynq or similar.

---

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

---

## 🙏 Acknowledgments

- [Echo Framework](https://echo.labstack.com/) - High performance, extensible, minimalist Go web framework
- [Go](https://golang.org/) - The Go programming language

---

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

_Made with ❤️ for the Go community_
