prod:
	@./build/tailwindcss-extra -c ./configs/tailwind.config.js -i ./web/css/input.css -o ./web/css/output.css -m
	@templ generate

dev:
	@./build/tailwindcss-extra -c ./configs/tailwind.config.js -i ./web/css/input.css -o ./web/css/output.css -m
	@templ generate
	@air -c ./scripts/.air.toml