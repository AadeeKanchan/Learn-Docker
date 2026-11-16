# Phase 5: Multi-Stage Builds & Image Optimization

Learn how to reduce container image sizes from hundreds of megabytes (or gigabytes) down to under 15 megabytes, dramatically speeding up deployments and eliminating security vulnerabilities (CVEs).

---

## 1. Why Multi-Stage Builds Matter

In traditional single-stage builds, everything needed to compile or package code (compilers, build tools, package managers, development headers, git, temporary build files) remains baked into the final image:
- Result: **Bloated image size (800MB - 2GB+)**, slower deployment transfers, and large attack surfaces full of potential CVEs.

**The Multi-Stage Solution**:
Use multiple `FROM` instructions in a single `Dockerfile`. Each `FROM` begins a new stage with a completely fresh environment. You selectively copy only the compiled binary or production artifacts from the builder stage into the final slim runtime stage!

```
+------------------------------------+        +-----------------------------+
|        STAGE 1: BUILDER            |        |      STAGE 2: FINAL         |
|  FROM golang:1.22-alpine AS builder |        |  FROM alpine:3.20 (or scratch)
|  - Compilers & build tools         |        |  - Zero compilers           |
|  - Full source code                |  COPY  |  - Zero source code         |
|  - Go module dependencies          | -----> |  - ONLY compiled executable |
|  - Compiles /app/server binary     |        |                             |
|  Size: ~350 MB                     |        |  Size: ~12 MB (or 5MB!)     |
+------------------------------------+        +-----------------------------+
```

---

## 2. Base Image Comparison Matrix

| Base Image | Typical Size | Includes Shell? | Includes Package Manager? | Best For |
|---|---|---|---|---|
| `ubuntu:24.04` / `debian:bookworm` | 80 MB - 120 MB | Yes (`bash`, `sh`) | Yes (`apt`) | Complex legacy apps, heavy system dependencies. |
| `python:slim` / `node:slim` | 40 MB - 70 MB | Yes (`sh`, `bash`) | Yes (`apt-get`) | Standard production web services. |
| `alpine:3.20` | ~5 MB - 7 MB | Yes (`/bin/sh`) | Yes (`apk`) | Extremely lightweight utilities, Go/Rust/Node runtimes (musl libc). |
| `gcr.io/distroless/static-debian12` | ~2 MB | **No** | **No** | Production hardened binaries; zero shell to exploit. |
| `scratch` | **0 Bytes** (Empty base) | **No** | **No** | Statically linked compiled binaries (Go, Rust, C++). |

---

## 3. Hands-On Lab: Shrinking a Service to 12 MB

Inspect the sample Go microservice in `sample-go-app/`.

### Step 1: Navigate to the directory
```powershell
cd 05-multi-stage-builds/sample-go-app
```

### Step 2: Build the multi-stage image
```powershell
docker build -t go-microservice:v1 .
```

### Step 3: Compare sizes!
Run:
```powershell
docker images go-microservice:v1
docker images golang:1.22-alpine
```
* Notice the full Go build environment is around **300 MB - 800 MB**.
* Your final production image (`go-microservice:v1`) is only **~12 MB**!

### Step 4: Run the optimized container
```powershell
docker run -d --name go-demo -p 8082:8080 go-microservice:v1
```

Test it in your browser or terminal:
```powershell
curl http://localhost:8082
curl http://localhost:8082/metrics
```

### Step 5: Cleanup
```powershell
docker stop go-demo
docker rm go-demo
cd ../..
```
