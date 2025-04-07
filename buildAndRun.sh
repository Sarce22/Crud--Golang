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
