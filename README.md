# Go Hexagonal Architecture

This project is an example of a Go application implementing **Hexagonal Architecture** (also known as Ports and Adapters). It isolates the core business logic from external concerns like databases and user interfaces.

## Project Structure

The project follows a standard layout for hexagonal applications:

- **`cmd/`**: Contains the main entry point of the application.
    - `api/main.go`: The bootstrapper that wires up dependencies and starts the server.
- **`internal/`**: Contains the application code, which is private to the project.
    - **`core/`**: The heart of the application.
        - `domain/`: Business entities and logic.
        - `ports/`: Interfaces (ports) that define interactions between the core and the outside world.
        - `services/`: Implementation of the core business logic (use cases).
    - **`adapters/`**: Implementations of the ports (adapters).
        - `handler/`: HTTP handlers (driving adapters).
        - `repository/`: Data storage implementations (driven adapters).

## Getting Started

### Prerequisites

- [Go](https://go.dev/) 1.25 or higher

### Running the Application

You can easily run the application using `make`:

```bash
make run
```

This will start the server on default port `8080`.

### Other Commands

- **Build functionality**: `make build`
- **Run tests**: `make test`
- **Tidy dependencies**: `make tidy`

## Usage & Testing

You can interact with the API using `curl`.

### Add a User

To create a new user, send a POST request:

```bash
curl -X POST http://localhost:8080/users \
     -H "Content-Type: application/json" \
     -d '{"name": "Edwin", "email": "edwin@codesandtags.io", "nickname": "edwin123"}'
```

### Get a User

To retrieve a user by their nickname:

```bash
curl -X GET http://localhost:8080/users/edwin123
```

## Testing & Coverage

Run unit tests:
```bash
go test -v ./...
```

Run tests with coverage:
```bash
go test -coverprofile=coverage.out ./...
```

View coverage report:
```bash
go tool cover -func=coverage.out
```
