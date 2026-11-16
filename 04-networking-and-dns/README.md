# Phase 4: Container Networking & Service Discovery

How containers communicate with each other, the host, and the outside world without IP address chaos.

---

## 1. Network Drivers Overview

| Driver | Description | Common Use Case |
|---|---|---|
| **`bridge` (Default)** | Creates an internal virtual network on a single host. Containers get private IPs and are isolated from the host LAN. | Default mode for standalone containers and multi-container apps on a single machine. |
| **`host`** | Bypasses Docker network virtualization. The container directly uses the host's network interfaces and ports. | Ultra-low latency, maximum network throughput (Note: Native to Linux, limited in Windows/macOS VM). |
| **`none`** | Completely disables networking. The container only has a loopback interface (`127.0.0.1`). | Secure batch jobs, offline computation, sensitive cryptographic operations. |
| **`overlay`** | Multi-host networking spanning across multiple Docker daemons/nodes. | Docker Swarm, clustered production microservices. |
| **`macvlan` / `ipvlan`**| Assigns a MAC/IP address to the container, making it appear as a physical device on the host's physical router/LAN. | Legacy apps requiring direct router subnets. |

---

## 2. Default Bridge vs. User-Defined Bridge (Critical Concept)

When you run a container without specifying `--network`, Docker attaches it to the **default bridge** network named `bridge`.

* **Default Bridge**:
  - ❌ **No automatic DNS resolution**: Containers cannot look up other containers by name (`ping container-b` fails!). You are forced to use unstable IP addresses or legacy `--link`.
  - ❌ Poor network isolation (all random containers share the same network).
* **User-Defined Bridge** (`docker network create my-net`):
  - ✅ **Automatic Embedded DNS Resolution**: Containers on the same custom network can communicate directly using their container name as the domain name (e.g., `http://auth-service:5000` or `redis://cache:6379`).
  - ✅ Strict security boundary: Containers on different user-defined networks cannot talk to each other unless explicitly connected.

---

## 3. Hands-On Lab: Multi-Container Service Discovery

We will build a two-tier application: A **Redis Cache** container and an interactive test client communicating purely via DNS.

### Step 1: Create a custom user-defined network
```powershell
docker network create my-app-network
```
Verify the network was created:
```powershell
docker network ls
```

### Step 2: Start a Redis container on the network
```powershell
docker run -d `
  --name redis-server `
  --network my-app-network `
  redis:alpine
```
Notice: We did **not** publish port 6379 to the host (`-p 6379:6379` is omitted). The database is securely hidden from external internet/host access, but accessible to any container on `my-app-network`!

### Step 3: Run an interactive container to test DNS resolution
```powershell
docker run -it --rm `
  --network my-app-network `
  redis:alpine sh
```

Inside the interactive shell:
```sh
# Ping the database using its container name!
ping -c 2 redis-server

# Use redis-cli to set and get a key over DNS:
redis-cli -h redis-server set my_key "Docker Networking is Awesome"
redis-cli -h redis-server get my_key

# Exit the shell
exit
```

### Step 4: Inspect the network
See the IP addresses and connected containers:
```powershell
docker network inspect my-app-network
```

### Step 5: Cleanup
```powershell
docker stop redis-server
docker rm redis-server
docker network rm my-app-network
```

---

## Self-Check Checklist
- [ ] Understand why user-defined bridge networks are preferred over the default bridge.
- [ ] Know how Docker's internal DNS resolver works.
- [ ] Know how to inspect connected containers with `docker network inspect`.
- [ ] Understand the security benefit of keeping databases unexposed to the host (`-p` omitted) while available internally.
