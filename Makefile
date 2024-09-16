prod:
	@sqlc generate -f ./configs/sqlc.yml
	@./scripts/tailwindcss-extra -c ./configs/tailwind.config.js -i ./web/css/input.css -o ./web/css/output.css -m
	@templ generate

dev:
	@sqlc generate -f ./configs/sqlc.yml
	@./scripts/tailwindcss-extra -c ./configs/tailwind.config.js -i ./web/css/input.css -o ./web/css/output.css -m
	@templ generate
	@air -c ./scripts/.air.toml

compose:
	@docker compose -f ./build/docker-compose.yml up -d --build
