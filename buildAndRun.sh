#!/bin/bash

# ============================
# Colores para salida legible
# ============================
GREEN='\033[0;32m'     # Verde
YELLOW='\033[1;33m'    # Amarillo
NC='\033[0m'           # Reset de color

# ============================
# Usuario de Docker Hub
# ============================
DOCKER_USER="sebastianarce"

# ============================
# Lista de microservicios
# ============================
SERVICES=("create" "read" "update" "delete")

# ============================
# Paso 1: Verificar cambios en el código
# ============================
echo -e "${GREEN}[INFO] Verificando cambios en los microservicios...${NC}"

CHANGED_SERVICES=()

# ============================
# Paso 2: Detectar cambios en cada microservicio con git status
# ============================
for SERVICE in "${SERVICES[@]}"; do
  if git status --porcelain | grep "${SERVICE}/" > /dev/null; then
    CHANGED_SERVICES+=("$SERVICE")
  fi
done

# ============================
# Paso 3: Si no hay cambios, salir
# ============================
if [ ${#CHANGED_SERVICES[@]} -eq 0 ]; then
  echo -e "${GREEN}[INFO] No se detectaron cambios en los microservicios.${NC}"
  exit 0
fi

# ============================
# Paso 4: Mostrar los microservicios con cambios
# ============================
echo -e "${YELLOW}[CAMBIOS DETECTADOS] Se modificaron:${NC}"
for svc in "${CHANGED_SERVICES[@]}"; do
  echo -e "  - $svc"
done

# ============================
# Paso 5: Confirmar si se desea continuar
# ============================
read -p "¿Deseas reconstruir las imágenes Docker y reiniciar los contenedores? (s/N): " confirm

# ============================
# Paso 6: Si se confirma, continuar
# ============================
if [[ "$confirm" =~ ^[sS]$ ]]; then

  # ============================
  # Paso 6.1: Preguntar si se desea subir a Git
  # ============================
  read -p "¿Deseas subir también los cambios a GitHub? (s/N): " push_git

  if [[ "$push_git" =~ ^[sS]$ ]]; then
    echo -e "${GREEN}[INFO] Agregando cambios a Git...${NC}"
    git add .

    # ============================
    # Paso 6.2: Pedir mensaje de commit
    # ============================
    read -p "Escribe un mensaje para el commit: " commit_message
    git commit -m "$commit_message"

    # ============================
    # Paso 6.3: Detectar rama actual y hacer push
    # ============================
    current_branch=$(git symbolic-ref --short HEAD)
    echo -e "${GREEN}[INFO] Haciendo push a la rama '${current_branch}'...${NC}"
    git push origin "$current_branch"
  else
    echo -e "${YELLOW}[GIT] Cambios locales no fueron subidos a GitHub.${NC}"
  fi

  # ============================
  # Paso 7: Build y push de imágenes Docker actualizadas
  # ============================
  for svc in "${CHANGED_SERVICES[@]}"; do
    IMAGE_NAME="${DOCKER_USER}/${svc}-golang:latest"

    echo -e "${GREEN}[DOCKER] Construyendo imagen: $IMAGE_NAME...${NC}"
    docker build -t $IMAGE_NAME ./$svc

    echo -e "${GREEN}[DOCKER] Subiendo imagen a Docker Hub: $IMAGE_NAME...${NC}"
    docker push $IMAGE_NAME
  done

  # ============================
  # Paso 8: Reiniciar contenedores
  # ============================
  echo -e "${GREEN}[INFO] Reiniciando contenedores...${NC}"
  docker-compose down
  docker-compose up -d

  echo -e "${GREEN}[✔ COMPLETADO] Código, imágenes y contenedores actualizados.${NC}"

else
  echo -e "${YELLOW}[CANCELADO] No se realizó ninguna acción.${NC}"
fi
