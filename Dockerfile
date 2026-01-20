# --------------------------------------------------------
# Stage 1: Builder
# Usamos la imagen oficial de Go para compilar
# --------------------------------------------------------
FROM golang:1.25-alpine AS builder

# Instalamos git por si alguna dependencia lo requiere
RUN apk add --no-cache git

WORKDIR /app

# Primero copiamos solo los archivos de dependencias.
# Esto aprovecha la caché de capas de Docker: si no cambias el go.mod,
# no se vuelven a descargar las librerías.
COPY go.mod go.sum ./
RUN go mod download

# Ahora copiamos el código fuente
COPY . .

# Compilamos el binario
# -o main: nombre del output
# cmd/api/main.go: la ruta de tu entrypoint
RUN go build -o main cmd/api/main.go

# --------------------------------------------------------
# Stage 2: Runner
# Usamos una imagen mínima para ejecutar
# --------------------------------------------------------
FROM alpine:latest

WORKDIR /root/

# Instalamos certificados CA por si tu app necesita hacer llamadas HTTPS externas
RUN apk --no-cache add ca-certificates
RUN apk add --no-cache git gcc musl-dev
# Habilitar CGO explícitamente al compilar
RUN CGO_ENABLED=1 go build -o main cmd/api/main.go

# Copiamos SOLO el binario desde la etapa "builder"
COPY --from=builder /app/main .

# Exponemos el puerto
EXPOSE 8080

# Comando de ejecución
CMD ["./main"]