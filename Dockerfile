# ── Build stage ──────────────────────────────────────────────
FROM golang:1.23-alpine AS builder

RUN apk add --no-cache ca-certificates

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -trimpath -o /logos ./cmd/logos

# ── Runtime stage (distroless) ───────────────────────────────
FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /logos /logos
COPY config.yaml /config.yaml

ENV CONFIG_PATH=/config.yaml

EXPOSE 3000

USER nonroot:nonroot

ENTRYPOINT ["/logos"]
