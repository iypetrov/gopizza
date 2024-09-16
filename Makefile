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
