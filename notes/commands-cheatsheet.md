# Docker Commands & Syntax Cheat Sheet

A comprehensive quick-reference for everyday Docker operations.

---

## 1. Container Lifecycle

| Command | Description | Example |
|---|---|---|
| `docker run` | Create and start a container | `docker run -d --name my-app -p 8080:80 nginx:alpine` |
| `docker start` | Start one or more stopped containers | `docker start my-app` |
| `docker stop` | Stop one or more running containers (SIGTERM then SIGKILL) | `docker stop my-app` |
| `docker restart` | Restart a container | `docker restart my-app` |
| `docker pause` / `unpause` | Suspend / resume all processes in a container | `docker pause my-app` |
| `docker kill` | Forcefully kill a container immediately (SIGKILL) | `docker kill my-app` |
| `docker rm` | Remove one or more stopped containers | `docker rm my-app` |
| `docker rm -f` | Force remove a running container | `docker rm -f my-app` |
| `docker rm $(docker ps -aq)` | Remove all stopped containers | (Run in PowerShell / bash) |

### Common `docker run` Flags:
- `-d`: Run container in detached mode (in the background).
- `-it`: Allocate pseudo-TTY and keep STDIN open (interactive shell).
- `--name <name>`: Assign a custom name to the container.
- `--rm`: Automatically remove the container when it exits.
- `-p <host_port>:<container_port>`: Publish a container port to the host.
- `-e KEY=VALUE`: Set environment variables inside the container.
- `--env-file <path>`: Load environment variables from a file.
- `-v <host_path>:<container_path>`: Bind mount a host directory or attach a named volume.
- `--network <net_name>`: Connect container to a specific network.
- `--restart <policy>`: Restart policy (`no`, `always`, `unless-stopped`, `on-failure`).
- `--cpus="1.5"`: Limit container to 1.5 CPU cores.
- `-m 512m`: Limit container memory to 512 Megabytes.

---

## 2. Container Inspection & Debugging

| Command | Description | Example |
|---|---|---|
| `docker ps` | List all running containers | `docker ps` |
| `docker ps -a` | List all containers (running and stopped) | `docker ps -a` |
| `docker logs` | Fetch logs of a container | `docker logs my-app` |
| `docker logs -f --tail 100` | Follow live container logs starting from last 100 lines | `docker logs -f --tail 100 my-app` |
| `docker exec -it` | Run a new command inside an active container | `docker exec -it my-app sh` |
| `docker top` | Display running processes of a container | `docker top my-app` |
| `docker stats` | Live stream of container CPU, RAM, Network I/O metrics | `docker stats` |
| `docker inspect` | Return low-level JSON details of a container or image | `docker inspect my-app` |
| `docker diff` | Inspect changes to files or directories on container filesystem | `docker diff my-app` |
| `docker port` | List port mappings for a container | `docker port my-app` |

---

## 3. Image Management

| Command | Description | Example |
|---|---|---|
| `docker build` | Build an image from a Dockerfile | `docker build -t my-username/app:1.0 .` |
| `docker build --no-cache` | Force build without using cached layers | `docker build --no-cache -t app:v1 .` |
| `docker images` / `docker image ls` | List locally stored images | `docker images` |
| `docker pull` | Pull an image from a registry | `docker pull redis:alpine` |
| `docker push` | Push an image to a registry | `docker push my-username/app:1.0` |
| `docker tag` | Create a target tag that references a source image | `docker tag app:v1 app:latest` |
| `docker rmi` / `docker image rm` | Remove one or more images | `docker rmi nginx:alpine` |
| `docker history` | Show history of layers for an image | `docker history my-app:1.0` |
| `docker save` | Save one or more images to a tar archive | `docker save -o app.tar my-app:1.0` |
| `docker load` | Load an image from a tar archive | `docker load -i app.tar` |

---

## 4. Storage & Volumes

| Command | Description | Example |
|---|---|---|
| `docker volume create` | Create a named volume | `docker volume create pg_data` |
| `docker volume ls` | List all volumes | `docker volume ls` |
| `docker volume inspect` | Display detailed information on a volume | `docker volume inspect pg_data` |
| `docker volume rm` | Remove one or more volumes | `docker volume rm pg_data` |
| `docker volume prune` | Remove all unused local volumes | `docker volume prune -f` |

---

## 5. Networking

| Command | Description | Example |
|---|---|---|
| `docker network create` | Create a custom user-defined network | `docker network create app-net` |
| `docker network ls` | List all networks | `docker network ls` |
| `docker network inspect` | Display detailed network info (IPs, connected containers) | `docker network inspect app-net` |
| `docker network connect` | Connect a running container to a network | `docker network connect app-net my-app` |
| `docker network disconnect`| Disconnect a container from a network | `docker network disconnect app-net my-app`|
| `docker network rm` | Remove one or more networks | `docker network rm app-net` |
| `docker network prune` | Remove all unused networks | `docker network prune -f` |

---

## 6. Docker Compose (V2)

| Command | Description | Example |
|---|---|---|
| `docker compose up -d` | Build, (re)create, and start all services detached | `docker compose up -d` |
| `docker compose down` | Stop and remove containers, networks created by up | `docker compose down` |
| `docker compose down -v` | Stop containers and also delete declared named volumes | `docker compose down -v` |
| `docker compose ps` | List status of services in the compose stack | `docker compose ps` |
| `docker compose logs -f` | Stream logs for all or specific services | `docker compose logs -f web` |
| `docker compose build` | Build or rebuild service images | `docker compose build --no-cache` |
| `docker compose restart` | Restart service containers | `docker compose restart api` |
| `docker compose exec` | Execute a command in a running service container | `docker compose exec db psql -U postgres`|

---

## 7. System Cleanup & Resource Management

| Command | Description |
|---|---|
| `docker system df` | Show Docker disk usage (images, containers, local volumes, build cache) |
| `docker system prune` | Remove unused data (stopped containers, unused networks, dangling images) |
| `docker system prune -a` | Deep clean: removes stopped containers, unused networks, and ALL images not associated with a container |
| `docker system prune -a --volumes` | Complete nuclear cleanup: frees maximum disk space (deletes all stopped containers, all unused images, and all unattached volumes) |
| `docker builder prune` | Clean up BuildKit build cache |
