# Stage 1: Build
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN apt-get update && apt-get install -y librdkafka-dev

ENV CGO_ENABLED=1

RUN go build -o main .

# Stage 2: Runtime
FROM debian:bookworm-slim

WORKDIR /app/

RUN apt-get update && apt-get install -y --no-install-recommends librdkafka1 ca-certificates && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/main .

EXPOSE 8080
CMD ["./main"]
