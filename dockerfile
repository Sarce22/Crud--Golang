# Etapa 1: Build
FROM golang:1.22.4-alpine AS builder

# Instala dependencias necesarias para compilar
RUN apk add --no-cache git

# Crea directorio de trabajo
WORKDIR /app

# Copia go.mod y go.sum primero para aprovechar la cache
COPY go.mod go.sum ./
RUN go mod download

# Copia el resto del código
COPY . .

# Compila el binario (binario estático para alpine)
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Etapa 2: Imagen liviana para producción
FROM alpine:3.19

# Crear usuario sin privilegios
RUN adduser -D appuser

WORKDIR /app

# Copiar el binario desde la etapa anterior
COPY --from=builder /app/main .

# Usar usuario sin privilegios
USER appuser

# Puerto expuesto (ajústalo si tu app usa otro)
EXPOSE 8080

# Comando por defecto
CMD ["./main"]
