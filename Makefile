prod:
	@./build/tailwindcss-extra -c ./configs/tailwind.config.js -i ./web/css/input.css -o ./web/css/output.css -m
	@templ generate

dev:
	@air -c ./scripts/.air.toml