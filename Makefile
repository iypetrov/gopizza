prod:
	@sqlc generate
	@./tailwindcss-extra -i ./web/css/input.css -o ./web/css/output.css -m
	@templ generate

dev:
	@sqlc generate
	@./tailwindcss-extra -i ./web/css/input.css -o ./web/css/output.css -m
	@templ generate
	@air -c .air.toml

compose:
	@docker compose up -d --build

fmt:	
	@go fmt ./...
	@templ fmt .
	@# sudo apt install -y pgformatter 
	@find . -name '*.sql' -exec pg_format -i {} +
