# Phase 2: Dockerfiles & Custom Image Creation

In this phase, you transition from running pre-built images to building your own container images from source code.

---

## 1. Complete Dockerfile Instruction Reference

A `Dockerfile` is a script containing sequential instructions that Docker reads to assemble an image.

| Instruction | Purpose | Best Practice / Common Pitfall |
|---|---|---|
| `FROM` | Sets the base image for subsequent instructions | Prefer official minimal tags (`python:3.12-slim`, `node:20-alpine`) over generic `:latest`. |
| `WORKDIR` | Sets the working directory inside the container | Always use absolute paths (`WORKDIR /app`). Creates the folder if it doesn't exist. Avoid multiple `RUN cd ...`. |
| `COPY` | Copies files from host build context into image | Prefer `COPY` over `ADD`. Only copy what's necessary (e.g. dependencies first to leverage caching). |
| `ADD` | Copies files; can also unpack local `.tar` files or fetch remote URLs | Only use `ADD` when you need automatic tar extraction. Use `curl`/`wget` with `RUN` for remote URLs. |
| `RUN` | Executes commands during image build time and commits results to a new layer | Chain related commands with `&& \` and clean package manager caches in the same layer (`apt-get clean && rm -rf /var/lib/apt/lists/*`). |
| `ENV` | Sets persistent environment variables (available during build and runtime) | Use for app configurations like `PORT=8000` or `NODE_ENV=production`. |
| `ARG` | Defines build-time variables (NOT available at runtime in the container) | Passed via `--build-arg KEY=value`. Do NOT use for production secrets. |
| `EXPOSE` | Documents which ports the container listens on at runtime | Acts as documentation/metadata. Does **not** publish the port to host on its own. |
| `VOLUME` | Creates a mount point and marks it as holding externally mounted volumes | Declares persistent storage locations. |
| `USER` | Sets the username or UID to use when running the image | Never run containers as `root` in production. Always switch to a non-root user. |
| `CMD` | Provides default arguments or command to execute upon container startup | Can be easily overridden at runtime via CLI arguments. |
| `ENTRYPOINT` | Configures a container that will run as an executable | Makes the container behave like a dedicated binary. Arguments passed via CLI are appended to `ENTRYPOINT`. |

---

## 2. `CMD` vs. `ENTRYPOINT` (The Definitive Comparison)

Both instructions define what process runs when the container starts, but they behave differently when CLI arguments are supplied.

### The Two Forms:
* **Exec form (Recommended)**: `["executable", "param1", "param2"]` (Runs directly without a shell wrapper; properly receives Unix signals like `SIGTERM`).
* **Shell form**: `command param1 param2` (Runs under `/bin/sh -c`, which means signals are sent to the shell, not your app process).

### Combining Them:
The industry best practice is using `ENTRYPOINT` for the fixed command and `CMD` for default arguments that can be overridden:

```dockerfile
ENTRYPOINT ["python", "app.py"]
CMD ["--port", "8000", "--workers", "2"]
```

- Running `docker run my-app`:
  Executes: `python app.py --port 8000 --workers 2`
- Running `docker run my-app --port 9090`:
  Executes: `python app.py --port 9090` (overrides the `CMD` default args).

---

## 3. The Layer Caching Strategy

Docker caches the result of each instruction. When you rebuild:
1. Docker checks if the instruction and its inputs have changed.
2. If unmodified, Docker uses the cached layer instantly.
3. **The moment a single layer changes, all subsequent layers are invalidated and must rebuild!**

### ❌ Inefficient Dockerfile (Breaks Cache Constantly)
```dockerfile
FROM python:3.12-slim
WORKDIR /app
COPY . .                            # Every small code edit invalidates this!
RUN pip install -r requirements.txt # Re-installs all pip packages every build!
CMD ["python", "app.py"]
```

### ✅ Optimized Dockerfile (Leverages Layer Caching)
```dockerfile
FROM python:3.12-slim
WORKDIR /app

# 1. Copy only dependency manifest first
COPY requirements.txt .

# 2. Install dependencies (Cached unless requirements.txt changes!)
RUN pip install --no-cache-dir -r requirements.txt

# 3. Copy application code last (changes frequently, but fast to copy)
COPY . .

CMD ["python", "app.py"]
```

---

## 4. Hands-On Lab: Containerizing a Python API

Inspect the sample application located in `sample-python-app/`.

### Step 1: Navigate to the sample directory
```powershell
cd 02-dockerfiles-and-images/sample-python-app
```

### Step 2: Build the Docker image
```powershell
docker build -t my-python-api:v1 .
```
Notice each step being executed and cached.

### Step 3: Run the container
```powershell
docker run -d --name python-demo -p 5000:5000 my-python-api:v1
```

### Step 4: Verify the application
Open your browser or terminal and test:
```powershell
curl http://localhost:5000
curl http://localhost:5000/health
```

### Step 5: Test Layer Caching
Make a minor edit to `app.py`, then rerun `docker build -t my-python-api:v2 .`.
Observe how Step 3 (`pip install`) finishes instantly with `CACHED`!

### Step 6: Cleanup
```powershell
docker stop python-demo
docker rm python-demo
cd ../..
```
