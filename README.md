# Chatapp Backend

A real-time chat REST API built with **Go**, **PostgreSQL**, **EMQX (MQTT)**, and **WebSocket**. Supports direct messages (DMs), group conversations, and live message delivery.

---

## Tech Stack

| Layer | Technology |
|---|---|
| Language | Go 1.25 |
| Framework | Gin |
| Database | PostgreSQL + GORM |
| Migrations | Goose |
| Auth | JWT (HS256) |
| Real-time | EMQX (MQTT) + WebSocket |
| IDs | Short nanoid (12 chars) |
| Containerization | Docker + Docker Compose |
| Orchestration | Kubernetes via k3d (K8s inside Docker) |

---

## Project Structure

```
chatapp-backend/
├── database/           # DB connection + Goose migrations
├── error/              # Centralized app error definitions
├── handler/            # HTTP layer (Gin handlers)
│   ├── auth/
│   ├── conversation/
│   └── user/
├── k8s/                # Kubernetes manifests
│   ├── namespace.yaml
│   ├── secret.yaml
│   ├── api/
│   ├── emqx/
│   └── postgres/
├── middleware/         # JWT auth middleware
├── models/             # Shared data models (DB + API)
├── mqtt/               # EMQX client + message notifier
├── repo/               # Database queries (GORM)
│   ├── auth/
│   ├── conversation/
│   └── user/
├── route/              # Route registration + dependency wiring
├── service/            # Business logic layer
│   ├── auth/
│   ├── conversation/
│   └── user/
├── utils/              # JWT helpers, password hashing, ID generation
├── ws/                 # WebSocket hub + client
├── main.go
├── Dockerfile
├── docker-compose.yml
└── Makefile
```

---

## Prerequisites

Choose one of the three ways to run the project below. Each has different prerequisites.

- **Local** — Go 1.21+, PostgreSQL, Goose, EMQX
- **Docker** — Docker + Docker Compose
- **Kubernetes** — Docker, k3d, kubectl

---

## Running the Project

### Option A — Local Development (fastest for coding)

Best when you're actively writing and testing code.

**1. Install dependencies**

```bash
git clone <repo-url>
cd chatapp-backend
go mod download
```

**2. Install Goose**

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

**3. Start EMQX (needed for real-time)**

```bash
docker run -d --name emqx \
  -p 1883:1883 -p 18083:18083 \
  emqx/emqx:latest
```

**4. Configure environment**

```bash
cp .env.example .env
# Edit .env with your Postgres credentials
```

**5. Run migrations and start**

```bash
make migrate-up
make run
```

Server: `http://localhost:8000`

---

### Option B — Docker Compose (everything in one command)

Best when you want Postgres + EMQX + API all running together without K8s.

**First time or after code changes:**
```bash
make docker-up
```

**Already running, no code changes:**
```bash
docker compose up -d
```

**Stop:**
```bash
make docker-down
```

**View logs:**
```bash
make docker-logs
```

| Service | URL |
|---------|-----|
| API | `http://localhost:8000` |
| EMQX Dashboard | `http://localhost:18083` |
| MQTT | `localhost:1883` |

---

### Option C — Kubernetes with k3d (production-like setup)

Best when you want to test the full deployment stack locally with Docker-based Kubernetes.

#### One-time setup

**1. Install k3d**
```bash
curl -s https://raw.githubusercontent.com/k3d-io/k3d/main/install.sh | bash
```

**2. Install kubectl**
```bash
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl
```

**3. Create cluster, build image, deploy**
```bash
make k8s-cluster   # creates K8s cluster inside Docker
make k8s-load      # builds image + imports into cluster
make k8s-deploy    # deploys Postgres + EMQX + API
```

**4. Verify everything is running**
```bash
make k8s-status
# All pods should show STATUS = Running
```

| Service | URL |
|---------|-----|
| API | `http://localhost:8000` |
| EMQX Dashboard | `http://localhost:30083` |
| MQTT | `localhost:1883` |

---

### Option D — Kubernetes with Minikube (VM-based Kubernetes)

Best when you prefer a VM-based Kubernetes environment or need more isolation.

#### One-time setup

**1. Install Minikube**
```bash
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube && rm minikube-linux-amd64
```

**2. Install kubectl** (if not already installed)
```bash
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl
```

**3. Start Minikube, build image, deploy**
```bash
make minikube-start   # starts Minikube cluster
make minikube-load    # builds image + loads into Minikube
make minikube-deploy  # deploys Postgres + EMQX + API
```

**4. Start Minikube tunnel** (in a separate terminal, keep it running)
```bash
make minikube-tunnel
# Or manually: minikube tunnel
# This allows you to access NodePort services on localhost
```

**5. Verify everything is running**
```bash
make minikube-status
# All pods should show STATUS = Running and READY 1/1 or 2/2
```

| Service | URL (after tunnel is running) |
|---------|-------------------------------|
| API | `http://localhost:8000` |
| EMQX Dashboard | `http://localhost:30083` |
| MQTT | `localhost:1883` |

---

#### Daily workflow comparison

**k3d:**
```bash
# Start: k3d cluster start chatapp
# Stop: k3d cluster stop chatapp
# After code changes: make k8s-redeploy
```

**Minikube:**
```bash
# Start: minikube start
# Stop: minikube stop
# After code changes: make minikube-redeploy
# Access services: make minikube-tunnel (separate terminal)
```

---

#### Every day workflow (k3d)
| API | `http://localhost:8000` |
| EMQX Dashboard | `http://localhost:30083` |
| MQTT | `localhost:1883` |

---

#### Every day workflow (K8s)

```bash
# Morning — start the cluster after PC restart
k3d cluster start chatapp

# Check pods came back up
make k8s-status

# Evening — stop to free RAM
k3d cluster stop chatapp
```

#### After changing code (K8s)

```bash
make k8s-load
# Rebuilds image + imports into cluster + restarts API pods automatically
```

#### Full reset (K8s)

```bash
make k8s-down      # delete cluster completely
make k8s-cluster   # recreate
make k8s-load
make k8s-deploy
```

---

## Quick Reference

| Situation | Command |
|-----------|---------|
| Run locally | `make run` |
| Start with Docker | `make docker-up` |
| Stop Docker | `make docker-down` |
| **k3d** | |
| Create k3d cluster (once) | `make k8s-cluster` |
| First deploy to k3d (once) | `make k8s-load && make k8s-deploy` |
| Start k3d after PC restart | `k3d cluster start chatapp` |
| Push code changes to k3d | `make k8s-redeploy` |
| Check k3d pod status | `make k8s-status` |
| View k3d API logs | `make k8s-logs` |
| Stop k3d for the day | `k3d cluster stop chatapp` |
| Full k3d wipe | `make k8s-down` |
| **Minikube** | |
| Start Minikube (first time) | `make minikube-start` |
| Start Minikube (after stop) | `minikube start` |
| First deploy to Minikube | `make minikube-load && make minikube-deploy` |
| Enable access to services | `make minikube-tunnel` (separate terminal) |
| Push code changes to Minikube | `make minikube-redeploy` |
| Check Minikube pod status | `make minikube-status` |
| View Minikube API logs | `make minikube-logs` |
| Stop Minikube for the day | `minikube stop` |
| Full Minikube wipe | `make minikube-delete` |

---

## Environment Variables

Copy `.env.example` to `.env` and fill in your values.

| Variable | Default | Description |
|----------|---------|-------------|
| `SERVER_PORT` | `8000` | HTTP server port |
| `POSTGRES_HOST` | `localhost` | Postgres host (`postgres` in Docker/K8s) |
| `POSTGRES_PORT` | `5432` | Postgres port |
| `POSTGRES_USER` | `postgres` | Postgres user |
| `POSTGRES_PASSWORD` | — | Postgres password |
| `POSTGRES_DB` | `chatapp` | Database name |
| `POSTGRES_SSLMODE` | `disable` | SSL mode |
| `JWT_SECRET` | — | Secret key for signing JWTs |
| `MQTT_BROKER` | `tcp://localhost:1883` | EMQX broker (`tcp://emqx:1883` in Docker/K8s) |
| `MQTT_USERNAME` | — | MQTT username (optional, leave empty for local) |
| `MQTT_PASSWORD` | — | MQTT password (optional, leave empty for local) |

> Docker Compose and K8s override `POSTGRES_HOST` and `MQTT_BROKER` automatically using their internal service names. Your `.env` stays unchanged for local dev.

---

## API Reference

All protected routes require:
```
Authorization: Bearer <token>
```

### Auth

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/signup` | Register a new user |
| POST | `/login` | Login and receive a JWT |

**POST /signup**
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "secret123"
}
```

**POST /login**
```json
{
  "email": "john@example.com",
  "password": "secret123"
}
```
Response includes a `token` — use it as your Bearer token for all protected routes.

---

### Users

| Method | Endpoint | Auth | Description |
|--------|----------|------|-------------|
| GET | `/users` | ✅ | List all registered users |

---

### Direct Messages (DMs)

| Method | Endpoint | Auth | Description |
|--------|----------|------|-------------|
| POST | `/dms` | ✅ | Start a DM (or retrieve existing one) |
| GET | `/dms` | ✅ | List all your DM conversations |
| POST | `/dms/:id/messages` | ✅ | Send a message |
| GET | `/dms/:id/messages` | ✅ | List messages |

**POST /dms**
```json
{ "user_id": "abc123xyz012" }
```

Response includes `other_user` so you immediately know who you're chatting with:
```json
{
  "conversation": {
    "id": "aB3xZ9mK2pLq",
    "type": "dm",
    "created_at": "...",
    "other_user": {
      "id": "xyz789abc012",
      "name": "Jane Smith",
      "email": "jane@example.com"
    }
  }
}
```

---

### Groups

| Method | Endpoint | Auth | Description |
|--------|----------|------|-------------|
| POST | `/groups` | ✅ | Create a group conversation |
| GET | `/groups` | ✅ | List all your group conversations |
| POST | `/groups/:id/messages` | ✅ | Send a message |
| GET | `/groups/:id/messages` | ✅ | List messages |

**POST /groups**
```json
{
  "name": "Team Alpha",
  "member_ids": ["abc123xyz012", "def456uvw789"]
}
```

---

### Message response shape

Every message includes the sender's name — no extra lookup needed on the frontend:

```json
{
  "id": "xZ9mK2aB3pLq",
  "conversation_id": "aB3xZ9mK2pLq",
  "sender_id": "abc123xyz012",
  "sender_name": "John Doe",
  "content": "Hey, how are you?",
  "created_at": "2026-07-30T12:00:00Z"
}
```

---

## Real-time with EMQX

When a message is sent via the REST API, the backend publishes it to EMQX on:

```
chat/conversation/{conversation_id}
```

### Test with MQTTX CLI

```bash
npm install -g mqttx-cli

# Subscribe to all conversations
mqttx sub -h localhost -p 1883 -t "chat/conversation/+" -v
```

### Test with MQTTX Desktop

1. New Connection → Host `localhost`, Port `1883`
2. Connect (no credentials needed for local EMQX)
3. Subscribe to `chat/conversation/+`
4. Send a message via the REST API — it appears instantly in MQTTX

### MQTT payload

```json
{
  "type": "message",
  "conversation_id": "aB3xZ9mK2pLq",
  "message": {
    "id": "xZ9mK2aB3pLq",
    "sender_id": "abc123xyz012",
    "sender_name": "John Doe",
    "content": "Hey!",
    "created_at": "2026-07-30T12:00:00Z"
  }
}
```

### EMQX Dashboard

| Mode | URL | Credentials |
|------|-----|-------------|
| Local / Docker | `http://localhost:18083` | `admin` / `public` |
| Kubernetes | `http://localhost:30083` | `admin` / `public` |

---

## WebSocket

Browser clients connect here for presence and real-time events:

```
ws://localhost:8000/ws?token=<jwt>
```

| Event type | Description |
|------------|-------------|
| `message` | New message in a conversation |
| `presence` | User came online or went offline |
| `online_users` | List of currently online user IDs (sent on connect) |

---

## Makefile Commands

```bash
# Local
make run                                  # Start the server
make migrate-up                           # Apply pending migrations
make migrate-down                         # Roll back last migration
make migrate-status                       # Show migration state
make create-migration name=my_migration   # Create a new migration file

# Docker
make docker-up                            # Build + start all services
make docker-down                          # Stop all services
make docker-logs                          # Tail API logs

# Kubernetes
make k8s-cluster                          # Create k3d cluster
make k8s-load                             # Build image + import + restart pods
make k8s-deploy                           # Apply all K8s manifests
make k8s-status                           # Show all pods and services
make k8s-logs                             # Tail API pod logs
make k8s-down                             # Delete the cluster
```
