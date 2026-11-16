# Phase 3: Storage & Data Persistence

Containers are ephemeral by default. When a container is deleted (`docker rm`), any data written inside its writable layer is permanently lost. This module teaches you how to persist data and sync local code with containers.

---

## 1. The Three Storage Mechanisms

```
+-------------------------------------------------------------+
|                          HOST OS                            |
|                                                             |
|  [ Bind Mount ]           [ Named Volume ]     [ tmpfs ]    |
|  C:\Projects\app          /var/lib/docker/     Host RAM     |
|          |                   volumes/my-vol         |       |
+----------|-------------------------|----------------|-------+
|          v                         v                v       |
|    /app/src (live dev)       /var/lib/postgresql  /tmp/cache|
|                                                             |
|                      CONTAINER                              |
+-------------------------------------------------------------+
```

| Type | Stored On Host | Managed By Docker | Best Use Case |
|---|---|---|---|
| **Named Volume** | Dedicated Docker storage (`/var/lib/docker/volumes/` inside Linux VM) | **Yes** (CLI commands: `docker volume ...`) | Databases (Postgres, MySQL, Mongo), stateful services, production data. |
| **Bind Mount** | Any arbitrary directory on your host file system (e.g. `C:\Users\...\app`) | **No** (Direct host file path access) | Local development with live source code reloading; sharing host config files. |
| **tmpfs Mount** | Host System Memory (RAM) only | **No** (Ephemeral memory) | Sensitive unencrypted secrets, high-speed temporary caching, secure keys. |

---

## 2. Syntax: `-v` vs. `--mount`

Docker supports two syntaxes:
1. `-v` or `--volume`: Older, shorter shorthand:
   - Named volume: `-v my_db_data:/var/lib/postgresql/data`
   - Bind mount: `-v ${PWD}/src:/app/src`
2. `--mount`: Modern, explicit syntax (recommended for production):
   - Named volume: `--mount type=volume,source=my_db_data,target=/var/lib/postgresql/data`
   - Bind mount: `--mount type=bind,source="C:\path\to\code",target=/app/code`
   - tmpfs mount: `--mount type=tmpfs,target=/app/cache`

---

## 3. Hands-On Labs

### Lab 1: Data Persistence with a Database (Named Volume)
Prove that data survives container deletion when using a named volume.

#### Step 1: Run PostgreSQL with a named volume
```powershell
docker run -d `
  --name db-demo `
  -e POSTGRES_PASSWORD=mysecretpassword `
  -v pgdata:/var/lib/postgresql/data `
  -p 5432:5432 `
  postgres:16-alpine
```

#### Step 2: Create a test table and insert data
```powershell
docker exec -it db-demo psql -U postgres -c "CREATE TABLE users (id SERIAL PRIMARY KEY, name VARCHAR(50));"
docker exec -it db-demo psql -U postgres -c "INSERT INTO users (name) VALUES ('Alice'), ('Bob');"
docker exec -it db-demo psql -U postgres -c "SELECT * FROM users;"
```

#### Step 3: Destroy the container
```powershell
docker stop db-demo
docker rm db-demo
```
The container is completely gone!

#### Step 4: Launch a brand new container attached to the same volume
```powershell
docker run -d `
  --name db-demo-restored `
  -e POSTGRES_PASSWORD=mysecretpassword `
  -v pgdata:/var/lib/postgresql/data `
  -p 5432:5432 `
  postgres:16-alpine
```

#### Step 5: Verify the data was preserved
```powershell
docker exec -it db-demo-restored psql -U postgres -c "SELECT * FROM users;"
```
Notice both Alice and Bob are still there!

#### Step 6: Cleanup
```powershell
docker stop db-demo-restored
docker rm db-demo-restored
docker volume rm pgdata
```

---

### Lab 2: Live Code Sync (Bind Mount)
Experience instant feedback without rebuilding the container image.

```powershell
# Create an empty local web folder
New-Item -ItemType Directory -Path "$PWD/html" -Force | Out-Null
Set-Content -Path "$PWD/html/index.html" -Value "<h1>Initial Version</h1>"

# Run Nginx with a bind mount to the local folder
docker run -d `
  --name live-web `
  -p 8081:80 `
  -v "${PWD}/html:/usr/share/nginx/html" `
  nginx:alpine
```

1. Open `http://localhost:8081` in your browser.
2. Edit `html/index.html` on your host machine (change it to `<h1>Instant Update without Rebuild!</h1>`).
3. Refresh `http://localhost:8081` in your browser. The page updates immediately!

#### Cleanup
```powershell
docker stop live-web
docker rm live-web
Remove-Item -Recurse -Force "$PWD/html"
```

---

### Lab 3: Backing Up and Restoring a Docker Volume
Learn the production-standard technique to backup any named volume to a `.tar.gz` archive on your host machine.

```powershell
# 1. Create a volume and write dummy data to it
docker volume create backup_test
docker run --rm -v backup_test:/data alpine sh -c "echo 'Important Database Data' > /data/backup.txt"

# 2. Backup the volume to a tar archive on your current host directory
docker run --rm -v backup_test:/volume -v "${PWD}:/backup" alpine tar -czf /backup/my-backup.tar.gz -C /volume .

# 3. Verify my-backup.tar.gz exists on your host
Get-ChildItem my-backup.tar.gz

# 4. Cleanup test volume and file
docker volume rm backup_test
Remove-Item my-backup.tar.gz
```

---

## Self-Check Checklist
- [ ] Understand why containers are ephemeral.
- [ ] Can differentiate between Bind Mounts, Named Volumes, and `tmpfs`.
- [ ] Can backup and restore named volume data.
- [ ] Understand how to safely prune unused volumes (`docker volume prune`).
