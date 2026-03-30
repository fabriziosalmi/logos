.PHONY: build run dev docker docker-up docker-down clean test

BINARY := logos
VERSION := $(shell cat VERSION)

build:
	CGO_ENABLED=0 go build -ldflags="-s -w" -trimpath -o $(BINARY) ./cmd/logos

run: build
	./$(BINARY)

dev:
	go run ./cmd/logos

docker:
	docker build -t logos:$(VERSION) -t logos:latest .

docker-up:
	docker compose up -d --build

docker-down:
	docker compose down

test:
	@echo "Starting server..."
	@go run ./cmd/logos & PID=$$!; \
	sleep 2; \
	echo "Testing endpoints..."; \
	curl -sf http://localhost:3000/healthz > /dev/null && echo "  [OK] /healthz"; \
	curl -sf http://localhost:3000/ > /dev/null && echo "  [OK] / (dashboard)"; \
	curl -sf http://localhost:3000/api/v4/render/favicon/pure/blue/spin.svg > /dev/null && echo "  [OK] V4 favicon"; \
	curl -sf http://localhost:3000/api/v4/render/og-card/spotlight/amber/cyan/pulse.svg > /dev/null && echo "  [OK] V4 og-card gradient"; \
	curl -sf http://localhost:3000/api/favicon/green/pulse.svg > /dev/null && echo "  [OK] V3 legacy"; \
	curl -sf http://localhost:3000/app/minerva/favicon.svg > /dev/null && echo "  [OK] App shortcut"; \
	echo "All tests passed!"; \
	kill $$PID 2>/dev/null

clean:
	rm -f $(BINARY)
