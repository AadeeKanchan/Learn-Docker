# ==========================================
# Stage 1: Build the Go executable
# ==========================================
FROM golang:1.22-alpine AS builder

WORKDIR /src

# Copy source code
COPY main.go .

# Compile a static binary without debug symbols for minimal size
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /bin/server main.go

# ==========================================
# Stage 2: Minimal Runtime Environment
# ==========================================
FROM alpine:3.20

# Create a non-privileged user
RUN adduser -D -u 1000 appuser
USER appuser

# Copy ONLY the static binary from the builder stage
COPY --from=builder /bin/server /bin/server

ENV PORT=8080
EXPOSE 8080

ENTRYPOINT ["/bin/server"]
