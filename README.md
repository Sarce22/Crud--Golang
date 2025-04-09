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
```

---

## 🔧 Script de despliegue automático

Este script `buildAndRun` permite automatizar la gestión del ciclo de vida de tus microservicios. Lo que hace es:

1. Verificar si hubo cambios en alguno de los microservicios (`create`, `read`, `update`, `delete`).
2. Pregunta si deseas reconstruir las imágenes Docker de los servicios modificados.
3. Opcionalmente, permite hacer commit y push de los cambios al repositorio Git.
4. Reconstruye las imágenes Docker afectadas y las sube a Docker Hub.
5. Reinicia el entorno con `docker compose` para aplicar los cambios.

Guarda este archivo como `buildAndRun` y dale permisos de ejecución (`chmod +x buildAndRun`). Aquí el contenido completo:

```bash
#!/bin/bash

# ============================
# Colores para salida legible
# ============================
GREEN='\033[0;32m'     # Verde
YELLOW='\033[1;33m'    # Amarillo
RED='\033[0;31m'       # Rojo
NC='\033[0m'           # Reset de color

# ============================
# Usuario de Docker Hub
# ============================
DOCKER_USER="sebastianarce"

# ============================
#  Lista de microservicios
# ============================
SERVICES=("create" "read" "update" "delete")

# ============================
# Paso 1: Verificar cambios
# ============================
echo -e "${GREEN}[INFO] Verificando cambios en los microservicios...${NC}"

CHANGED_SERVICES=()

for SERVICE in "${SERVICES[@]}"; do
  if git status --porcelain | grep "${SERVICE}/" > /dev/null; then
    CHANGED_SERVICES+=("$SERVICE")
  fi
done

# ============================
# Paso 2: Salir si no hay cambios
# ============================
if [ ${#CHANGED_SERVICES[@]} -eq 0 ]; then
  echo -e "${GREEN}[INFO] No se detectaron cambios en los microservicios.${NC}"
  exit 0
fi

# ============================
# Paso 3: Mostrar cambios
# ============================
echo -e "${YELLOW}[CAMBIOS DETECTADOS] Se modificaron:${NC}"
for svc in "${CHANGED_SERVICES[@]}"; do
  echo -e "  - $svc"
done

# ============================
# Paso 4: Confirmación del usuario
# ============================
read -p "¿Deseas reconstruir las imágenes Docker y reiniciar los contenedores? (s/N): " confirm

if [[ "$confirm" =~ ^[sS]$ ]]; then

  # ============================
  # Paso 5: Push opcional a GitHub
  # ============================
  read -p "¿Deseas subir también los cambios a GitHub? (s/N): " push_git

  if [[ "$push_git" =~ ^[sS]$ ]]; then
    echo -e "${GREEN}[GIT] Agregando cambios...${NC}"
    git add .

    read -p "📝 Escribe un mensaje para el commit: " commit_message
    git commit -m "$commit_message"

    current_branch=$(git symbolic-ref --short HEAD)
    echo -e "${GREEN}[GIT] Haciendo push a '${current_branch}'...${NC}"
    git push origin "$current_branch"
  else
    echo -e "${YELLOW}[GIT] Cambios locales NO fueron subidos a GitHub.${NC}"
  fi

  # ============================
  # Paso 6: Build y push Docker
  # ============================
  for svc in "${CHANGED_SERVICES[@]}"; do
    IMAGE_NAME="${DOCKER_USER}/${svc}-golang:latest"

    echo -e "${GREEN}[DOCKER] Construyendo imagen: $IMAGE_NAME...${NC}"
    docker build -t "$IMAGE_NAME" "./$svc" || { echo -e "${RED}[ERROR] Falló el build de $svc${NC}"; exit 1; }

    echo -e "${GREEN}[DOCKER] Subiendo imagen a Docker Hub: $IMAGE_NAME...${NC}"
    docker push "$IMAGE_NAME" || { echo -e "${RED}[ERROR] Falló el push de $svc${NC}"; exit 1; }
  done

  # ============================
  # Paso 7: Reiniciar contenedores
  # ============================
  echo -e "${GREEN}[INFO] Reiniciando contenedores...${NC}"
  docker compose down
  docker compose up -d --build

  echo -e "${GREEN}[✔ COMPLETADO] Código, imágenes y contenedores actualizados.${NC}"

else
  echo -e "${YELLOW}[CANCELADO] No se realizó ninguna acción.${NC}"
fi
```

---

