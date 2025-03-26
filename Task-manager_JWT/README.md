# Task Manager API with MongoDB

A RESTful API for managing tasks, built with Go and MongoDB.

## Prerequisites

- Go 1.16 or higher
- MongoDB 4.4 or higher
- Git

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/task-manager.git
cd task-manager
```

2. Install dependencies:
```bash
go mod download
```

3. Start MongoDB:
- Windows: Start MongoDB service
- Linux/Mac: `sudo service mongod start` or `brew services start mongodb-community`

4. Configure environment variables (optional):
```bash
# Windows (PowerShell)
$env:MONGODB_URI="mongodb://localhost:27017"
$env:PORT="8080"

# Linux/Mac
export MONGODB_URI="mongodb://localhost:27017"
export PORT="8080"
```

## Running the Application

1. Start the server:
```bash
go run main.go
```

2. The server will start on `http://localhost:8080`

## API Documentation

See [API Documentation](docs/API.md) for detailed endpoint information and examples.

## Project Structure

```
.
├── controllers/     # Request handlers
├── data/           # Database operations
├── docs/           # Documentation
├── errors/         # Custom error types
├── models/         # Data models
├── router/         # Route definitions
├── main.go         # Application entry point
└── go.mod          # Go module file
```

## MongoDB Configuration

The application uses the following MongoDB settings:
- Database: `taskmanager`
- Collection: `tasks`
- Default connection string: `mongodb://localhost:27017`

To use a different MongoDB instance:
1. Set the `MONGODB_URI` environment variable
2. Format: `mongodb://username:password@host:port`

## Error Handling

The API implements comprehensive error handling for:
- Database connection issues
- Invalid input data
- Resource not found
- Validation errors

## Development

### Adding New Features
1. Create new models in `models/`
2. Add database operations in `data/`
3. Create controllers in `controllers/`
4. Define routes in `router/`

### Testing
```bash
go test ./...
```

## License

MIT License 