build:
	@./tailwindcss-extra -i ./static/css/input.css -o ./static/css/output.css
	@templ generate

dev:
	@air -c .air.toml