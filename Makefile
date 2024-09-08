APP_NAME = gopizza

build:
	@echo "database..."
	@make queries
	@echo "frontend..."
	@npm ci
	@make frontend
	@echo "application..."
	@go build -o bin/main cmd/$(APP_NAME)/main.go

fmt:
	@gofmt -w -s .

queries:
	@sqlc generate

frontend:
	@npx tailwindcss build -i static/css/style.css -o static/css/tailwind.css -m
	@templ generate

dev:
	@npx tailwindcss -i static/css/style.css -o static/css/tailwind.css --watch \
	 &templ generate -watch -proxy=http://localhost:8080 -open-browser=false

compose:
	@docker compose up -d --build
