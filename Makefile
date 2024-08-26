APP_NAME = gopizza

build:
	@go build -o bin/main cmd/$(APP_NAME)/main.go

dev:
	@docker compose up -d
