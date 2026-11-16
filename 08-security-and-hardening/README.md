# Phase 8: Container Security & Hardening

By default, Docker containers run as the Linux `root` user (UID 0). If an attacker manages to break out of a container via a kernel exploit or misconfiguration, they inherit root privileges on your host machine. This phase teaches you how to lock down containers to enterprise security standards.

---

## 1. The Core Hardening Principles

```
+-----------------------------------------------------------+
|              DEFENSE IN DEPTH STRATEGY                    |
|                                                           |
|  1. Non-Root User        -> Run as unprivileged UID 1000  |
|  2. Read-Only Rootfs     -> Attacker cannot write malware |
|  3. Drop Capabilities    -> Strip kernel system calls     |
|  4. No New Privileges    -> Block setuid escalation       |
|  5. Scanned Images       -> Zero known CVEs               |
+-----------------------------------------------------------+
```

### Principle 1: Never Run as `root`
Always create and switch to a dedicated user in your `Dockerfile`:
```dockerfile
RUN adduser -D -u 1000 appuser && chown -R appuser:appuser /app
USER appuser
```

### Principle 2: Read-Only Root Filesystem (`--read-only`)
Prevent attackers from modifying application code, downloading malicious scripts, or tampering with binaries:
```powershell
docker run -d --read-only --name secure-app my-app:latest
```
If your application *must* write temporary logs or cache files, mount an ephemeral memory `tmpfs` specifically to that directory:
```powershell
docker run -d --read-only --tmpfs /tmp --tmpfs /var/run -p 80:80 nginx:alpine
```

### Principle 3: Dropping Linux Capabilities
Linux breaks root privileges into granular capabilities (e.g. `CAP_NET_RAW` to forge network packets, `CAP_SYS_ADMIN` to mount filesystems). Most web applications need none of them:
```powershell
# Drop ALL Linux kernel capabilities, then add back only what is required:
docker run -d `
  --cap-drop=ALL `
  --cap-add=NET_BIND_SERVICE `
  -p 80:80 `
  nginx:alpine
```

### Principle 4: Prevent Privilege Escalation
Ensure processes cannot gain more privileges than their parent process:
```powershell
--security-opt=no-new-privileges:true
```

---

## 2. Vulnerability Scanning with Docker Scout & Trivy

Before pushing images to production, always scan them for Common Vulnerabilities and Exposures (CVEs):

### Using Docker Scout:
```powershell
# Quick summary of vulnerabilities
docker scout quickview <image-name>

# Detailed list of CVEs with remediation advice
docker scout cves <image-name>

# Compare vulnerabilities between two image tags
docker scout compare --to <image-name>:latest <image-name>:v2
```

### Using Trivy (Open Source):
```powershell
# Run Trivy directly inside an ephemeral container to scan another image!
docker run --rm -v /var/run/docker.sock:/var/run/docker.sock aquasec/trivy:latest image my-app:v1
```

---

## 3. Hands-On Lab: Assembling a Fully Hardened Container

Let's run a bulletproof Nginx container that has:
1. All Linux capabilities dropped.
2. Read-only root filesystem.
3. No privilege escalation.
4. Ephemeral `tmpfs` mounts for temporary cache and runtime PIDs.

```powershell
docker run -d `
  --name hardened-nginx `
  -p 8085:80 `
  --read-only `
  --tmpfs /var/cache/nginx:rw,noexec,nosuid `
  --tmpfs /var/run:rw,noexec,nosuid `
  --tmpfs /tmp:rw,noexec,nosuid `
  --cap-drop=ALL `
  --cap-add=CHOWN `
  --cap-add=SETGID `
  --cap-add=SETUID `
  --cap-add=NET_BIND_SERVICE `
  --security-opt=no-new-privileges:true `
  nginx:alpine
```

### Test the security boundary:
Attempt to create a file inside the container:
```powershell
docker exec -it hardened-nginx touch /malicious_file
```
**Result**: `touch: /malicious_file: Read-only file system`!

### Cleanup:
```powershell
docker stop hardened-nginx
docker rm hardened-nginx
```

---

## Self-Check Checklist
- [ ] Understand why running containers as root is dangerous.
- [ ] Can implement `--read-only` root filesystems combined with `--tmpfs`.
- [ ] Understand Linux capabilities and how to `--cap-drop=ALL`.
- [ ] Know how to run image vulnerability scans with `docker scout`.
