# Docker Mastery: 0 to Advanced Hands-on Labs

Welcome to your structured hands-on Docker learning workspace. This repository is organized into progressive, modular labs designed to take you from container foundations to production-grade orchestration and security hardening.

[Docker Notes](https://drive.google.com/file/d/1pNl761VoBkFG8M15xGfrrsRH-uQssc5y/view?usp=sharing)

---

## 🗂️ Learning Curriculum & Module Index

| Module | Title | Core Focus | Link |
|---|---|---|---|
| **01** | Foundations & Essential CLI | Container vs VM, Namespaces, Cgroups, UnionFS, `run`, `ps`, `exec`, `logs`, port mapping | [Phase 1 Guide](file:///c:/Users/adika/OneDrive/Desktop/Docker/01-foundations-and-cli/README.md) |
| **02** | Dockerfiles & Image Creation | Instructions (`FROM`, `WORKDIR`, `COPY`, `RUN`, `CMD`), layer caching, Python API sample | [Phase 2 Guide](file:///c:/Users/adika/OneDrive/Desktop/Docker/02-dockerfiles-and-images/README.md) |
| **03** | Storage & Data Persistence | Bind mounts, named volumes, `tmpfs`, database state, volume backup/restore | [Phase 3 Guide](file:///c:/Users/adika/OneDrive/Desktop/Docker/03-storage-and-volumes/README.md) |
| **04** | Networking & Service Discovery | Bridge networks, embedded DNS, container-to-container communication, port exposure | [Phase 4 Guide](file:///c:/Users/adika/OneDrive/Desktop/Docker/04-networking-and-dns/README.md) |
| **05** | Multi-Stage Builds & Optimization | Builder patterns, Alpine vs Distroless vs Scratch, shrinking images from 800MB to 12MB | [Phase 5 Guide](file:///c:/Users/adika/OneDrive/Desktop/Docker/05-multi-stage-builds/README.md) |
| **06** | Docker Compose Orchestration | Multi-tier stacks, `compose.yaml`, healthchecks, environment configs, fullstack demo | [Phase 6 Guide](file:///c:/Users/adika/OneDrive/Desktop/Docker/06-docker-compose/README.md) |
| **07** | Production Operations & Ops | Resource limits (CPU/RAM), OOM killer (code 137), restart policies, log rotation, prune | [Phase 7 Guide](file:///c:/Users/adika/OneDrive/Desktop/Docker/07-production-and-ops/README.md) |
| **08** | Security & Hardening | Non-root users (`USER`), `--read-only`, dropping Linux capabilities, vulnerability scans | [Phase 8 Guide](file:///c:/Users/adika/OneDrive/Desktop/Docker/08-security-and-hardening/README.md) |
| **09** | BuildKit, Multi-Arch & CI/CD | Docker Buildx, cross-platform (`amd64`/`arm64`), cache mounts, GitHub Actions pipelines | [Phase 9 Guide](file:///c:/Users/adika/OneDrive/Desktop/Docker/09-advanced-buildx-cicd/README.md) |
| **Notes** | Commands & Syntax Cheat Sheet | Quick reference of all CLI commands, flags, and shortcuts | [Cheat Sheet](file:///c:/Users/adika/OneDrive/Desktop/Docker/notes/commands-cheatsheet.md) |

---

## 🚀 Quick Verification: Day 1 Check

Ensure Docker Desktop is open on your Windows machine, then run:

```powershell
# 1. Check version and connectivity
docker version

# 2. Run the test container
docker run hello-world

# 3. Spin up an Nginx web server
docker run -d --name test-nginx -p 8080:80 nginx:alpine
```
Visit `http://localhost:8080` in your web browser. When verified:
```powershell
docker stop test-nginx
docker rm test-nginx
```

---

## 📌 Progress Tracker
- [ ] **Phase 1**: Foundations & CLI Lifecycle
- [ ] **Phase 2**: Dockerfiles & Image Authoring
- [ ] **Phase 3**: Storage, Bind Mounts & Named Volumes
- [ ] **Phase 4**: Networking & Service Discovery
- [ ] **Phase 5**: Multi-Stage Builds & Size Optimization
- [ ] **Phase 6**: Multi-Container Docker Compose
- [ ] **Phase 7**: Production Operations & Resource Limits
- [ ] **Phase 8**: Container Security & Hardening
- [ ] **Phase 9**: BuildKit, Multi-Arch & CI/CD Pipelines
