# Phase 7: Production Operations, Monitoring & Resource Governance

Running containers in development is forgiving; running them in 24/7 production environments requires guardrails against memory leaks, runaway CPU usage, crashing processes, and disk exhaustion.

---

## 1. Restart Policies

What should Docker do when your container crashes or the host reboots?

| Policy | Behavior | Production Recommendation |
|---|---|---|
| `no` (Default) | Do not automatically restart the container under any circumstance. | Good for batch jobs / migration scripts. |
| `on-failure[:max-retries]` | Restarts only if the container exits with a non-zero exit code (error). | Good for scripts that should retry on temporary failure. |
| `always` | Always restart the container if it stops. If stopped manually, it restarts when the daemon reboots! | Avoid in favor of `unless-stopped`. |
| `unless-stopped` | Always restarts regardless of exit code, unless explicitly stopped by a human operator (`docker stop`). | **Gold standard for production web services & databases.** |

---

## 2. Resource Limits & The OOM (Out of Memory) Killer

Without resource constraints, a single rogue container with a memory leak can consume 100% of host RAM, causing the Linux Kernel to freeze or crash the host machine!

### Key Flags:
* `-m 512m` or `--memory="512m"`: Hard memory limit. If the container exceeds this, the Linux kernel **OOM Killer** immediately sends `SIGKILL` (Exit Code `137`).
* `--memory-reservation="256m"`: Soft limit; Docker warns or restricts memory only when the host begins running low on memory.
* `--cpus="1.5"`: Guarantees the container gets at most 1.5 CPU cores of computing power.
* `--cpu-shares=512`: Relative priority when multiple containers compete for CPU cycles (default is 1024).

### In `compose.yaml`:
```yaml
services:
  web:
    image: my-app:latest
    deploy:
      resources:
        limits:
          cpus: "1.0"
          memory: 512M
        reservations:
          cpus: "0.25"
          memory: 128M
```

---

## 3. Health Checks

By default, Docker only checks if the main container process (PID 1) is running. If your web app enters a deadlock or memory hang and stops answering HTTP requests, Docker still reports it as "Running" (UP).

### Adding a `HEALTHCHECK`:
```dockerfile
HEALTHCHECK --interval=30s --timeout=3s --start-period=10s --retries=3 \
  CMD curl -f http://localhost:8080/health || exit 1
```
* `interval`: Frequency of health checks.
* `timeout`: Maximum time to wait for a response before counting as failure.
* `start-period`: Grace period during container bootstrap before failures count toward retries.
* `retries`: Consecutive failures required to transition container to `unhealthy`.

---

## 4. Logging Drivers & Preventing Disk Exhaustion

By default, Docker captures STDOUT/STDERR using the `json-file` driver. Without configuration, **these log files grow infinitely until your hard drive runs out of disk space!**

### Configuring Log Rotation:
Configure rotation either globally in `/etc/docker/daemon.json` (or Docker Desktop Settings) or per service:

```yaml
services:
  api:
    image: my-api:latest
    logging:
      driver: "json-file"
      options:
        max-size: "10m"   # Rotate when file reaches 10 Megabytes
        max-file: "3"     # Keep at most 3 historical log files
```

---

## 5. Hands-On Lab: Observing the OOM Killer (Exit Code 137)

Prove that memory limits protect your host by intentionally exceeding a container's limit.

### Step 1: Run a container capped at only 20 MB of RAM
```powershell
docker run --name oom-test -m 20m alpine sh -c "python3 -c 'a = [0] * 100000000'" 2>$null
```
*(Or using a shell command to allocate memory beyond 20MB)*

### Step 2: Check the container exit code
```powershell
docker ps -a --filter "name=oom-test"
```
Notice:
- Exit code is **137** (`128 + 9 (SIGKILL)`).
- Run `docker inspect oom-test --format '{{.State.OOMKilled}}'`: returns `true`!

### Step 3: Cleanup
```powershell
docker rm oom-test
```

---

## 6. Housekeeping & Maintenance Cheat Sheet
```powershell
# See breakdown of disk space consumed by images, containers, and build cache:
docker system df

# Clean up dangling resources safely:
docker system prune

# Clean up EVERYTHING unused (all stopped containers, unused networks, and unused images):
docker system prune -a --volumes
```
