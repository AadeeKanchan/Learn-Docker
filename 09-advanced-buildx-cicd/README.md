# Phase 9: Advanced BuildKit, Multi-Architecture Builds & CI/CD

In modern enterprise environments, code is compiled for multiple hardware architectures (Intel/AMD x86_64, ARM64 for Apple Silicon & AWS Graviton) and automated via Continuous Integration / Continuous Deployment (CI/CD) pipelines.

---

## 1. Next-Generation BuildKit Features

Ensure BuildKit is active (default in modern Docker Desktop):
```powershell
$env:DOCKER_BUILDKIT=1
```

### 1. Cache Mounts (`--mount=type=cache`)
Eliminate re-downloading packages every time dependencies change. BuildKit preserves a persistent cache across builds:
```dockerfile
# For Python pip packages
RUN --mount=type=cache,target=/root/.cache/pip \
    pip install -r requirements.txt

# For Node.js npm packages
RUN --mount=type=cache,target=/root/.npm \
    npm ci

# For Go modules
RUN --mount=type=cache,target=/go/pkg/mod \
    go build -o /bin/app .
```

### 2. Secret Mounts (`--mount=type=secret`)
Pass credentials (like private SSH keys or API tokens) into a build **without leaving traces in intermediate layers or the final image**:
```dockerfile
RUN --mount=type=secret,id=my_secret \
    TOKEN=$(cat /run/secrets/my_secret) && curl -H "Authorization: Bearer $TOKEN" ...
```
Build command:
```powershell
docker build --secret id=my_secret,src=token.txt -t app:v1 .
```

---

## 2. Multi-Architecture Builds with Docker Buildx

Why do multi-architecture builds matter?
- Developers frequently work on **ARM64** laptops (Apple Silicon M-series, ARM Windows).
- Production servers frequently run on **AMD64 / x86_64** (Intel/AMD) or **ARM64** (AWS Graviton, Ampere).
- An image built only on AMD64 will fail or run with severe QEMU performance penalties on ARM64!

### Setting Up a Multi-Architecture Builder:
```powershell
# 1. Create a new buildx builder instance using the docker-container driver
docker buildx create --name my-builder --driver docker-container --bootstrap --use

# 2. Verify supported platforms (e.g. linux/amd64, linux/arm64)
docker buildx inspect my-builder
```

### Building & Pushing Multi-Arch Images:
```powershell
docker buildx build `
  --platform linux/amd64,linux/arm64 `
  -t your-username/my-service:1.0.0 `
  --push .
```
Docker automatically creates an **OCI Image Index (Manifest List)**. When a client pulls `your-username/my-service:1.0.0`, Docker automatically serves the exact binary compiled for that machine's CPU architecture!

---

## 3. Production GitHub Actions CI/CD Workflow Template

Place this in your repository at `.github/workflows/docker-build-push.yml` to automatically build, scan, and publish images whenever you push code:

```yaml
name: Build and Publish Docker Image

on:
  push:
    branches: [ "main" ]
    tags: [ 'v*.*.*' ]
  pull_request:
    branches: [ "main" ]

env:
  REGISTRY: ghcr.io
  IMAGE_NAME: ${{ github.repository }}

jobs:
  build-and-push:
    runs-on: ubuntu-latest
    permissions:
      contents: read
      packages: write

    steps:
      - name: Checkout repository
        uses: actions/checkout@v4

      # Set up QEMU for multi-architecture emulation (amd64, arm64)
      - name: Set up QEMU
        uses: docker/setup-qemu-action@v3

      # Set up Docker Buildx
      - name: Set up Docker Buildx
        uses: docker/setup-buildx-action@v3

      # Authenticate to GitHub Container Registry
      - name: Log in to Registry
        if: github.event_name != 'pull_request'
        uses: docker/login-action@v3
        with:
          registry: ${{ env.REGISTRY }}
          username: ${{ github.actor }}
          password: ${{ secrets.GITHUB_TOKEN }}

      # Extract metadata (tags, labels)
      - name: Extract Docker metadata
        id: meta
        uses: docker/metadata-action@v5
        with:
          images: ${{ env.REGISTRY }}/${{ env.IMAGE_NAME }}

      # Build and push with layer caching enabled
      - name: Build and push Docker image
        uses: docker/build-push-action@v5
        with:
          context: .
          platforms: linux/amd64,linux/arm64
          push: ${{ github.event_name != 'pull_request' }}
          tags: ${{ steps.meta.outputs.tags }}
          labels: ${{ steps.meta.outputs.labels }}
          cache-from: type=gha
          cache-to: type=gha,mode=max
```

---

## Self-Check Checklist
- [ ] Understand how BuildKit cache mounts speed up package installations.
- [ ] Understand why secret mounts are safer than `ARG` or `ENV` for build credentials.
- [ ] Know how to create a Buildx builder and compile multi-arch images (`linux/amd64`, `linux/arm64`).
- [ ] Understand how GitHub Actions automates image building and pushing to registries.
