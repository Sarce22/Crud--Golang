# 🚀 CRUD Microservicio en Go + MongoDB + Docker

Microservicio RESTful modular desarrollado en **Golang**, con base de datos **MongoDB** y desplegado usando **Docker**. Gestiona usuarios con operaciones **CRUD** completas (Create, Read, Update, Delete).

---

## 🧩 Características principales

- 📦 Arquitectura limpia: controlador, servicio y repositorio.
- 🧠 Validación de campos obligatorios.
- 🛡️ Prevención de duplicados por cédula y correo electrónico.
- 🧪 Pruebas unitarias con `testify`.
- 🐳 Contenedores Docker optimizados.
- ☁️ Imágenes subidas a Docker Hub.

---


---

## ⚙️ docker-compose.yml

Puedes utilizar el siguiente archivo `docker-compose.yml` para levantar todo el entorno:

```yaml
version: '3.8'

services:
  mongo:
    image: mongo:latest
    container_name: mongoDB
    ports:
      - "27017:27017"
    environment:
      MONGO_INITDB_DATABASE: testdb
    volumes:
      - mongodb-data:/data/db
      - mongodb-backup:/backup
    healthcheck:
      test: echo 'db.runCommand("ping").ok' | mongosh mongo:27017/test --quiet
      interval: 10s
      retries: 5
    restart: unless-stopped
    networks:
      - app-network

  create-service:
    image: sebastianarce/create-golang
    container_name: create-golang
    ports:
      - "8080:8080"
    depends_on:
      mongo:
        condition: service_healthy
    environment:
      MONGO_URI: mongodb://mongo:27017
      MONGO_DATABASE: testdb
    restart: unless-stopped
    networks:
      - app-network

  read-service:
    image: sebastianarce/read-golang
    container_name: read-golang
    ports:
      - "8082:8080"
    depends_on:
      mongo:
        condition: service_healthy
    environment:
      MONGO_URI: mongodb://mongo:27017
      MONGO_DATABASE: testdb
    restart: unless-stopped
    networks:
      - app-network

  update-service:
    image: sebastianarce/update-golang
    container_name: update-golang
    ports:
      - "8083:8080"
    depends_on:
      mongo:
        condition: service_healthy
    environment:
      MONGO_URI: mongodb://mongo:27017
      MONGO_DATABASE: testdb
    restart: unless-stopped
    networks:
      - app-network

  delete-service:
    image: sebastianarce/delete-golang
    container_name: delete-golang
    ports:
      - "8081:8080"
    depends_on:
      mongo:
        condition: service_healthy
    environment:
      MONGO_URI: mongodb://mongo:27017
      MONGO_DATABASE: testdb
    restart: unless-stopped
    networks:
      - app-network

volumes:
  mongodb-data:
  mongodb-backup:

networks:
  app-network:


