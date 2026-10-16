# BLKP DIY Periode 4 2026 - Golang Backend Engineering for Industries

## 🚀 Project Overview

This is a **Go Gin Framework** REST API backend built for industry engineering purposes. The application provides a clean, modular structure for managing todo items with full CRUD operations, validation, and PostgreSQL integration.

## 📁 Project Structure

```
.
├── main.go                 # Entry point, starts the server
├── router/                 # Router configuration
│   └── router.go           # Sets up all routes under /api/v1
├── todo/                   # Todo management module
│   ├── handler.go          # Business logic handlers
│   ├── routes.go           # Route registration
│   └── todos.go            # Todo data model
├── database/               # Database layer
│   └── databases.go        # DB connection setup
├── go.mod                  # Go module definition
└── .env                    # Environment variables (do NOT commit)
```

## 🛠️ Features

- **REST API** with Gin Framework
- **CRUD Operations** for Todo items (Create, Read, Update, Delete)
- **PostgreSQL** database connectivity (using pgx driver)
- **Input Validation** with go-playground/validator
- **Error Handling** with consistent JSON responses
- **Modular Architecture** - clean separation of concerns

## 📋 API Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | `/api/v1` | List all todos (with optional `?done=true/false` filter) |
| GET | `/api/v1/:id` | Get a specific todo by ID |
| POST | `/api/v1` | Create a new todo |
| PUT | `/api/v1/:id` | Update a todo |
| DELETE | `/api/v1/:id` | Delete a todo |
| GET | `/health` | Health check endpoint |

## 🔧 Getting Started

### Prerequisites

- Go 1.27+
- PostgreSQL database
- Docker (optional, for local development)

### Installation

1. Clone the repository
2. Set up environment variables in `.env`:
   ```env
   DATABASE_URL=postgresql://postgres:ferianLEO123@localhost:5432/postgres
   ```
3. Run the application:
   ```bash
   go mod tidy
   go run .
   ```

### Running Tests

```bash
go test ./...
```

### Building

```bash
go build
```

## 🧪 Testing

The project includes handlers with validation. You can add unit tests in `todo/handler_test.go` for comprehensive coverage.

## 📦 Dependencies

- **Gin Framework** (`github.com/gin-gonic/gin`) - Web framework
- **Gonix** (`github.com/jackc/pgx/v5/stdlib`) - PostgreSQL driver
- **Go Playground Validator** (`github.com/go-playground/validator/v10`) - Input validation
- **Godotenv** (`github.com/joho/godotenv`) - Load environment variables

## 🐛 Common Issues

- **Database connection**: Ensure `DATABASE_URL` is set correctly
- **Port conflict**: Make sure port 8080 is available
- **Missing env vars**: Check `.env` file is present and valid

## 📄 License

This project is licensed under the MIT License.

## 👥 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/your-feature`)
3. Commit your changes (`git commit -m 'Add feature'`)
4. Push to the branch (`git push origin feature/your-feature`)
5. Open a Pull Request

## 📞 Support

For questions, refer to the project documentation or contact the developer.