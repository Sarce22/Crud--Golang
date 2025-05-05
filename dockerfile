# Etapa 1: Imagen base de construcción
FROM golang:1.20-alpine as builder

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar go.mod y go.sum primero, para que los paquetes de dependencias se descarguen más rápido
COPY go.mod go.sum ./

# Descargar las dependencias
RUN go mod download

# Copiar el código fuente completo
COPY . .

# Establecer las variables de entorno de Go
ENV GO111MODULE=on

# Construir el binario optimizado
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o main .

# Etapa 2: Imagen final ligera
FROM alpine:latest

# Establecer el directorio de trabajo para la etapa final
WORKDIR /root/

# Copiar el binario desde la etapa anterior
COPY --from=builder /app/main .

# Exponer el puerto si es necesario
EXPOSE 8080

# Ejecutar el binario
CMD ["./main"]
