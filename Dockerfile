# ── Stage 1: Build ───────────────────────────────────────────────────────────
FROM golang:1.23-alpine AS builder

# Install build tools
RUN apk add --no-cache git

# GOTOOLCHAIN=auto lets Go automatically download the correct toolchain
# version declared in go.mod (1.25.7) even though the base image has 1.23
ENV GOTOOLCHAIN=auto

WORKDIR /app

# Download dependencies first (cached layer — only re-runs if go.mod/go.sum change)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source and build a static binary
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o chatapp-backend .

# ── Stage 2: Run ─────────────────────────────────────────────────────────────
FROM alpine:3.20

# ca-certificates needed for HTTPS/TLS connections (EMQX TLS, etc.)
RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

# Copy only the compiled binary and migration files
COPY --from=builder /app/chatapp-backend .
COPY --from=builder /app/database/migration ./database/migration

EXPOSE 8000

CMD ["./chatapp-backend"]
