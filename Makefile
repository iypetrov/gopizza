prod:
	@sqlc generate
	@./tailwindcss-extra -i ./public/css/input.css -o ./public/css/output.css --minify
	@templ generate
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags prod -o bin/main main.go static_prod.go

local-linux:
	@docker ps -a --format '{{.Names}}' | grep -wq gopizza_db || (docker volume create gopizza_data && docker run -d --name gopizza_db --network host -e POSTGRES_DB=gopizza -e POSTGRES_USER=user -e POSTGRES_PASSWORD=pass -v gopizza_data:/var/lib/postgresql/data postgres:15)
	@sqlc generate
	@stripe listen --forward-to http://localhost:8080/api/v0/payments/webhook --forward-connect-to http://localhost:8080/api/v0/payments/webhook & \
	./tailwindcss-extra -i ./public/css/input.css -o ./public/css/output.css --minify --watch & \
	templ generate --watch --proxy="http://localhost:8080" --open-browser=false & \
	air -c .air.toml

local-mac:
	@docker ps -a --format '{{.Names}}' | grep -wq gopizza_db || (docker volume create gopizza_data && docker run -d --name gopizza_db --network host -e POSTGRES_DB=gopizza -e POSTGRES_USER=user -e POSTGRES_PASSWORD=pass -v gopizza_data:/var/lib/postgresql/data postgres:15)
	@sqlc generate
	@./tailwindcss-extra-macos -i ./public/css/input.css -o ./public/css/output.css --minify
	@templ generate
	@stripe listen --forward-to http://localhost:8080/api/v0/payments/webhook --forward-connect-to http://localhost:8080/api/v0/payments/webhook & \
	air -c .air.toml

fmt:	
	@go fmt ./...
	@goimports -l -w .
	@templ fmt .
	@find . -name '*.sql' -exec pg_format -i {} +
