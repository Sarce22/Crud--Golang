# Etapa de construcción
FROM golang:1.22.7-alpine AS builder

WORKDIR /app

# Instalar dependencias necesarias
RUN apk add --no-cache git

# Copiar el código fuente
COPY . .

# Descargar dependencias y compilar el binario
RUN go mod tidy && go build -o main .

# Etapa de ejecución mínima
FROM alpine:3.19

# Copiar binario compilado desde la etapa anterior
COPY --from=builder /app/main /app/main

# Establecer el punto de entrada
ENTRYPOINT ["/app/main"]
