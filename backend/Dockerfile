# ──────────────────────────────────────────────────────────────────────
# Zplus Calendar – Backend Dockerfile (multi-stage)
# ──────────────────────────────────────────────────────────────────────

# ── Stage 1: Build ──────────────────────────────────────────────────
FROM golang:1.22-alpine AS builder

RUN apk add --no-cache git ca-certificates tzdata

WORKDIR /build

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source and build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-w -s" -o /build/bin/zplus-calendar ./cmd/server

# ── Stage 2: Runtime ───────────────────────────────────────────────
FROM alpine:3.19

RUN apk add --no-cache ca-certificates tzdata curl

# Non-root user
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

# Copy binary and config
COPY --from=builder /build/bin/zplus-calendar .
COPY --from=builder /build/configs ./configs

# Create uploads directory
RUN mkdir -p /app/uploads && chown -R appuser:appgroup /app

USER appuser

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=5s --start-period=10s --retries=3 \
    CMD curl -f http://localhost:8080/api/v1/health || exit 1

ENTRYPOINT ["./zplus-calendar"]
