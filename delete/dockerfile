# Etapa 1: Compilación
FROM golang:1.20-alpine AS builder
WORKDIR /app

# Copiar dependencias primero para aprovechar caché
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código fuente
COPY . .

# Compilar el binario optimizado
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o main .

# Etapa 2: Imagen final ligera
FROM alpine:latest
WORKDIR /root/

# Copiar solo el binario desde la etapa de compilación
COPY --from=builder /app/main .

# Exponer el puerto y ejecutar el binario
EXPOSE 8080
CMD ["./main"]
