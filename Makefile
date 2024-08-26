APP_NAME = gopizza

build:
	@go build -o bin/main cmd/$(APP_NAME)/main.go

local:
	@air -c .air.toml

dev:
	@docker compose up -d
