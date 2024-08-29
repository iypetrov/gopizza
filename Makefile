APP_NAME = gopizza

build:
	@sqlc generate
	@go build -o bin/main cmd/$(APP_NAME)/main.go

run:
	@sqlc generate
	@air -c .air.toml

local:
	@docker compose up -d

migrate_local:
	@goose -dir sql/migrations postgres "postgresql://user:pass@localhost:5432/gopizza?sslmode=disable" up