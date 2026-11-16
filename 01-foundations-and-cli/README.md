# Phase 1: Docker Foundations & Essential CLI

Welcome to the first phase of your Docker mastery journey. This module covers the theoretical mental model of containers and gives you hands-on experience with core container lifecycle commands.

---

## 1. Mental Model: Containers vs. Virtual Machines

### The Fundamental Difference
* **Virtual Machines (VMs)**: Each VM includes a full Guest Operating System (hundreds of MBs to several GBs), running on top of a Hypervisor (Type 1 or Type 2). Each VM emulates physical hardware. Boot time is measured in minutes, and CPU/RAM overhead is substantial.
* **Containers**: Containers share the **Host OS Linux Kernel**. They are simply isolated standard operating system processes running in user space. Boot time is instantaneous (milliseconds), and resource consumption is minimal.

```
       [ Virtual Machines ]                    [ Containers ]
+---------------------------------+     +--------------------------+
|  App 1   |  App 2   |  App 3    |     | App 1  |  App 2 | App 3  |
+----------+----------+-----------+     +--------+--------+--------+
| Bins/Libs| Bins/Libs| Bins/Libs |     | Bins/Libs (User space)   |
+----------+----------+-----------+     +--------------------------+
| Guest OS | Guest OS | Guest OS  |     |      Docker Engine       |
+----------+----------+-----------+     +--------------------------+
|           Hypervisor            |     |     Host Linux Kernel    |
+---------------------------------+     +--------------------------+
|        Host Operating System    |     |      Host Hardware       |
+---------------------------------+     +--------------------------+
|          Host Hardware          |
+---------------------------------+
```

### The Three Linux Primitives Powering Containers
Docker is not "magic"—it is an abstraction layer built on top of three Linux kernel features:
1. **Namespaces (Isolation)**: Provides what a process can *see*.
   - `pid`: Process IDs (process inside container thinks it is PID 1).
   - `net`: Network devices, ports, routing tables.
   - `mnt`: Mount points and file systems.
   - `ipc`: Inter-process communication resources.
   - `uts`: Hostname and domain name.
   - `user`: User and group IDs.
2. **Control Groups / Cgroups (Resource Metering)**: Enforces what a process can *use* (maximum CPU cores, RAM limits, swap space, I/O rates).
3. **Union File System (UnionFS / Overlay2)**: Allows directories and files from multiple image layers to be combined into a single coherent unified view using **Copy-on-Write (CoW)**.

---

## 2. Image vs. Container

* **Image (Class / Blueprint)**: An immutable, read-only package containing application code, runtimes, system tools, libraries, and default configurations.
* **Container (Object / Instance)**: A runnable, isolated process instantiated from an image. A container adds a thin, writable layer on top of the underlying read-only image layers.

---

## 3. Hands-On Labs

### Lab 1: First Flight (`hello-world`)
Test that your Docker Engine is connected and operational.

```powershell
docker run hello-world
```
**What happens under the hood?**
1. Docker client contacts the local Docker daemon.
2. The daemon checks local image cache for `hello-world`. Not finding it, it pulls it from Docker Hub.
3. The daemon creates a new container from that image, runs its greeting binary which prints to STDOUT, and exits.

---

### Lab 2: Interactive Linux Shell
Explore a real Ubuntu container environment interactively.

```powershell
docker run -it --rm ubuntu:latest bash
```
* **Flags explained**:
  - `-i`: Interactive mode (keeps standard input open).
  - `-t`: Allocates a pseudo-TTY terminal.
  - `--rm`: Automatically deletes the container filesystem when you type `exit`.

Inside the container shell, test process isolation:
```bash
# Check PID 1 (it's your bash shell, not Windows systemd/init!)
ps aux

# Check the hostname (it's the container's generated ID)
hostname

# Check OS release
cat /etc/os-release

# Exit when done
exit
```

---

### Lab 3: Running a Detached Web Server with Port Mapping
Run an Nginx web server in the background and expose it to your browser.

```powershell
docker run -d --name my-web-server -p 8080:80 nginx:alpine
```
* **Flags explained**:
  - `-d`: Detached mode (runs container in the background, freeing your terminal).
  - `--name my-web-server`: Gives the container an easily referenceable custom name.
  - `-p 8080:80`: Maps host port `8080` to container port `80`.

**Verify**:
1. Open your browser and navigate to `http://localhost:8080`. You will see the "Welcome to nginx!" screen.
2. Check running status:
   ```powershell
   docker ps
   ```
3. Inspect live logs:
   ```powershell
   docker logs -f my-web-server
   ```
   (Refresh your browser page, and watch the HTTP GET requests stream in real-time. Press `Ctrl+C` to stop following logs.)

---

### Lab 4: Executing Commands Inside a Running Container
Modify files inside your running Nginx container without rebuilding the image:

```powershell
# Open an interactive shell inside the running container
docker exec -it my-web-server sh

# Change the homepage HTML
echo "<h1>Hello from Docker Inside My Web Server!</h1>" > /usr/share/nginx/html/index.html

# Exit container shell
exit
```
Now refresh `http://localhost:8080` in your browser. Notice the custom message!

---

### Lab 5: Cleanup & Housekeeping
Always learn to clean up after testing:

```powershell
# Stop the web server
docker stop my-web-server

# Delete the container
docker rm my-web-server

# Verify it is gone
docker ps -a
```

---

## Self-Check Checklist
- [ ] Can explain the difference between a Container and a VM.
- [ ] Understand the role of Namespaces, Cgroups, and Overlay2.
- [ ] Understand the difference between `-d`, `-it`, `-p`, `--name`, and `--rm`.
- [ ] Comfortable inspecting logs with `docker logs -f` and executing shell commands with `docker exec -it`.
- [ ] Successfully tested Nginx and accessed `http://localhost:8080`.
