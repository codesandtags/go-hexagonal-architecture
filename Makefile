# Run the application
run:
	@go run cmd/api/main.go

# Build the application
build:
	@go build -o bin/api cmd/api/main.go

# Run tests
test:
	@go test ./... -v

# Tidy dependencies
tidy:
	@go mod tidy

# Run tests with coverage
test-cover:
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out

# Build Docker image
docker-build:
	@docker build -t go-hexagonal .

# Run Docker container
docker-run:
	@docker run -p 8080:8080 go-hexagonal ./main -db=sqlite

# Stop Docker container
docker-stop:
	@docker stop go-hexagonal