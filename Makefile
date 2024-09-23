prod:
	@sqlc generate
	@./tailwindcss-extra -i ./web/css/input.css -o ./web/css/output.css --minify
	@templ generate
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/main cmd/gopizza/main.go

dev:
	@docker compose up -d --build
	@sqlc generate
	@./tailwindcss-extra -i ./web/css/input.css -o ./web/css/output.css --minify --watch & \
	templ generate --watch --proxy="http://localhost:8080" --open-browser=false

fmt:	
	@go fmt ./...
	@templ fmt .
	@find . -name '*.sql' -exec pg_format -i {} +
