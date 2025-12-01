# ============================
#  BUILD STAGE
# ============================
FROM golang:1.24.5 AS builder

WORKDIR /app

# Cache deps
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build server
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o bin/server ./cmd/server

# Build worker
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o bin/worker ./cmd/worker

# ============================
#  RUNTIME STAGE
# ============================
FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/bin/server .
COPY --from=builder /app/bin/worker .

RUN addgroup -S app && adduser -S app -G app
USER app

CMD ["./server"]
