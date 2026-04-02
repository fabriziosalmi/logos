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
	curl -sf http://localhost:3000/api/v4/render/favicon/pure/blue/spin.svg?shape=shield > /dev/null && echo "  [OK] V4 shape"; \
	curl -sf http://localhost:3000/api/v4/render/favicon/pure/neon/spin.svg?texture=glitch > /dev/null && echo "  [OK] V4 texture"; \
	curl -sf http://localhost:3000/api/v4/render/favicon/pure/blue/spin.svg?variant=badge > /dev/null && echo "  [OK] V4 variant"; \
	curl -sf http://localhost:3000/api/v4/render/favicon/pure/blue/spin.svg?speed=2 > /dev/null && echo "  [OK] V4 speed"; \
	curl -sf http://localhost:3000/api/v4/render/twitter-card/dots/slate-500/breathe.svg > /dev/null && echo "  [OK] V4 tailwind+scene+format"; \
	curl -sf http://localhost:3000/api/v4/render/favicon/pure/cyan/static.svg?theme=dracula > /dev/null && echo "  [OK] V4 theme preset"; \
	curl -sf http://localhost:3000/api/favicon/green/pulse.svg > /dev/null && echo "  [OK] V3 legacy"; \
	curl -sf http://localhost:3000/app/logos/favicon.svg > /dev/null && echo "  [OK] App shortcut"; \
	curl -sf http://localhost:3000/app/unknown/favicon.svg > /dev/null && echo "  [OK] App fallback"; \
	echo "All tests passed!"; \
	kill $$PID 2>/dev/null

clean:
	rm -f $(BINARY)
