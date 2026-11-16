# Phase 6: Multi-Container Orchestration with Docker Compose

Running multiple containers individually with long, complex `docker run` commands becomes unmanageable for real applications. Docker Compose allows you to define and run entire multi-service architectures with a single declarative configuration file.

---

## 1. The Anatomy of `compose.yaml`

Modern Docker Compose uses the Compose V2 specification (prefer `compose.yaml` over legacy `docker-compose.yml`):

```yaml
services:
  # Service 1: Web Frontend
  frontend:
    image: nginx:alpine
    ports:
      - "80:80"
    volumes:
      - ./web:/usr/share/nginx/html:ro
    depends_on:
      api:
        condition: service_healthy
    networks:
      - app-tier

  # Service 2: Backend API
  api:
    build:
      context: ./api
      dockerfile: Dockerfile
    environment:
      - DATABASE_URL=postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@database:5432/${POSTGRES_DB}
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:5000/health"]
      interval: 10s
      timeout: 5s
      retries: 3
    networks:
      - app-tier
      - data-tier

  # Service 3: Database
  database:
    image: postgres:16-alpine
    environment:
      POSTGRES_USER: ${POSTGRES_USER}
      POSTGRES_PASSWORD: ${POSTGRES_PASSWORD}
      POSTGRES_DB: ${POSTGRES_DB}
    volumes:
      - db-data:/var/lib/postgresql/data
    networks:
      - data-tier

volumes:
  db-data:

networks:
  app-tier:
  data-tier:
```

---

## 2. Essential Docker Compose CLI Commands

| Command | Action |
|---|---|
| `docker compose up` | Creates networks/volumes, builds images, and starts all containers in foreground |
| `docker compose up -d` | Starts the entire stack in the background (detached mode) |
| `docker compose up -d --build` | Rebuilds images before starting (ensures fresh code changes take effect) |
| `docker compose down` | Stops and tears down containers, internal networks, and default links |
| `docker compose down -v` | **Deletes declared named volumes** as well (resets state completely) |
| `docker compose ps` | Displays live status of all services in the stack |
| `docker compose logs -f [service]` | Streams combined or service-specific logs |
| `docker compose exec [service] [cmd]` | Runs an interactive command inside a running service container |
| `docker compose restart [service]` | Restarts a specific service container |
| `docker compose top` | Displays running processes across the whole stack |

---

## 3. Hands-On Lab: Launching a Full Stack

Explore the sample stack in `sample-stack/`.

### Step 1: Navigate to the stack folder
```powershell
cd 06-docker-compose/sample-stack
```

### Step 2: Copy the sample environment file
```powershell
Copy-Item .env.example .env
```

### Step 3: Launch the stack
```powershell
docker compose up -d
```
Notice how Docker Compose automatically:
1. Creates the custom isolated bridge networks.
2. Creates the persistent named volume.
3. Pulls and configures the images.
4. Starts the services in correct dependency order.

### Step 4: Verify the stack status
```powershell
docker compose ps
```

### Step 5: Test the application
Open your browser and navigate to:
```
http://localhost:8080
```

### Step 6: Stream real-time logs
```powershell
docker compose logs -f
```
(Press `Ctrl+C` to stop watching logs.)

### Step 7: Teardown cleanly
```powershell
# Stop services while preserving database volume:
docker compose down

# Or stop and completely delete the database volume:
docker compose down -v
cd ../..
```

---

## Self-Check Checklist
- [ ] Understand the structure of a `compose.yaml` file (services, volumes, networks).
- [ ] Understand how `depends_on` and health checks coordinate startup order.
- [ ] Know how environment variable interpolation (`.env`) works.
- [ ] Comfortable managing multi-container stacks with `docker compose up -d` and `docker compose down`.
