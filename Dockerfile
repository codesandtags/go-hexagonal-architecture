# --------------------------------------------------------
# Stage 1: Builder
# Oficial Go image
# --------------------------------------------------------
FROM golang:1.25-alpine AS builder

# We need git for some dependencies
RUN apk add --no-cache git gcc musl-dev

WORKDIR /app

# First we copy only the dependency files.
# This takes advantage of Docker's layer cache: if you don't change go.mod,
# the libraries won't be downloaded again.
COPY go.mod go.sum ./
RUN go mod download

# Now we copy the source code
COPY . .

# Enable CGO explicitly when compiling
RUN CGO_ENABLED=1 go build -o main cmd/api/main.go

# --------------------------------------------------------
# Stage 2: Runner
# Usamos una imagen mínima para ejecutar
# --------------------------------------------------------
FROM alpine:latest

WORKDIR /root/

# We install CA certificates in case your app needs to make external HTTPS calls
RUN apk --no-cache add ca-certificates

# We copy ONLY the binary from the "builder" stage
COPY --from=builder /app/main .

# Environment variables
ENV DB_TYPE=memory

# We expose the port
EXPOSE 8080

# Comando de ejecución
CMD ["./main"]