# Etapa 1: Construcción
FROM golang:1.20-alpine AS builder

# Crear y establecer el directorio de trabajo
WORKDIR /app

# Copiar los archivos mod y sum primero para aprovechar el cache
COPY go.mod go.sum ./

# Descargar las dependencias
RUN go mod download

# Copiar todo el código fuente del proyecto
COPY . .

# Construir el binario optimizado
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o main .

# Etapa 2: Imagen final ligera
FROM alpine:latest

# Crear el directorio de trabajo en la imagen final
WORKDIR /root/

# Copiar el binario desde la etapa de construcción
COPY --from=builder /app/main .

# Exponer el puerto que usa la aplicación
EXPOSE 8080

# Ejecutar el binario
CMD ["./main"]
