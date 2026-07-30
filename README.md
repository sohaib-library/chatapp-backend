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
└── Makefile
```

---

## Prerequisites

- Go 1.21+
- PostgreSQL
- [Goose](https://github.com/pressly/goose) — `go install github.com/pressly/goose/v3/cmd/goose@latest`
- EMQX broker — easiest via Docker:

```bash
docker run -d --name emqx \
  -p 1883:1883 \
  -p 8083:8083 \
  -p 18083:18083 \
  emqx/emqx:latest
```

---

## Setup

**1. Clone and install dependencies**

```bash
git clone <repo-url>
cd chatapp-backend
go mod download
```

**2. Configure environment**

Copy the example and fill in your values:

```bash
cp .env.example .env
```

```env
SERVER_PORT=8000

POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=chatapp
POSTGRES_SSLMODE=disable

JWT_SECRET=change-me

# EMQX MQTT broker (anonymous by default — no username/password needed for local dev)
MQTT_BROKER=tcp://localhost:1883
MQTT_USERNAME=
MQTT_PASSWORD=
```

**3. Run migrations**

```bash
make migrate-up
```

**4. Start the server**

```bash
make run
# or
go run main.go
```

Server starts on `http://localhost:8000`

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
Response includes a `token` field — use this as your Bearer token.

---

### Users

| Method | Endpoint | Auth | Description |
|--------|----------|------|-------------|
| GET | `/users` | ✅ | List all registered users |

---

### Direct Messages (DMs)

| Method | Endpoint | Auth | Description |
|--------|----------|------|-------------|
| POST | `/dms` | ✅ | Start a DM with another user |
| GET | `/dms` | ✅ | List all your DM conversations |
| POST | `/dms/:id/messages` | ✅ | Send a message in a DM |
| GET | `/dms/:id/messages` | ✅ | List messages in a DM |

**POST /dms** — start or retrieve existing DM
```json
{ "user_id": "abc123xyz012" }
```
Response includes `other_user` with the name and ID of the person you're chatting with:
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
| POST | `/groups/:id/messages` | ✅ | Send a message in a group |
| GET | `/groups/:id/messages` | ✅ | List messages in a group |

**POST /groups**
```json
{
  "name": "Team Alpha",
  "member_ids": ["abc123xyz012", "def456uvw789"]
}
```

---

### Messages response shape

Every message includes the sender's name — no extra lookup needed:

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

When a message is sent via the REST API, the backend **publishes it to EMQX** on topic:

```
chat/conversation/{conversation_id}
```

### Subscribe with MQTTX CLI

```bash
# Install
npm install -g mqttx-cli

# Subscribe to all conversations
mqttx sub -h localhost -p 1883 -t "chat/conversation/+" -v
```

### Subscribe with MQTTX Desktop

1. Open MQTTX → New Connection
2. Host: `localhost`, Port: `1883`
3. Connect (no username/password needed for local EMQX)
4. Add subscription: `chat/conversation/+`

### MQTT message payload

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

View connected clients, topics, and live messages at:
```
http://localhost:18083
```
Default credentials: `admin` / `public`

---

## WebSocket

Browser clients can connect to the WebSocket endpoint for presence and real-time events:

```
ws://localhost:8000/ws?token=<jwt>
```

Event types received over WebSocket:

| Type | Description |
|------|-------------|
| `message` | New message in a conversation |
| `presence` | User came online or went offline |
| `online_users` | List of currently online user IDs (sent on connect) |

---

## Makefile Commands

```bash
make run              # Start the server
make migrate-up       # Apply all pending migrations
make migrate-down     # Roll back the last migration
make migrate-status   # Show migration status
make create-migration name=your_migration_name  # Create a new migration file
```

---

## Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `SERVER_PORT` | `8000` | HTTP server port |
| `POSTGRES_HOST` | `localhost` | Postgres host |
| `POSTGRES_PORT` | `5432` | Postgres port |
| `POSTGRES_USER` | `postgres` | Postgres user |
| `POSTGRES_PASSWORD` | — | Postgres password |
| `POSTGRES_DB` | `chatapp` | Database name |
| `POSTGRES_SSLMODE` | `disable` | SSL mode |
| `JWT_SECRET` | — | Secret key for signing JWTs |
| `MQTT_BROKER` | `tcp://localhost:1883` | EMQX broker address |
| `MQTT_USERNAME` | — | MQTT username (optional) |
| `MQTT_PASSWORD` | — | MQTT password (optional) |
