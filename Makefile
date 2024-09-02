APP_NAME = gopizza

build:
	@sqlc generate
	@go build -o bin/main cmd/$(APP_NAME)/main.go

fmt:
	@gofmt -w -s .

run:
	@sqlc generate
	@air -c .air.toml

local:
	@docker compose up -d --build
