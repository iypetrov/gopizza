FROM golang:1.23.1 AS build-stage
WORKDIR /app
COPY . .
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
RUN go install github.com/a-h/templ/cmd/templ@latest
RUN make prod

FROM gcr.io/distroless/base-debian12 AS run-stage
COPY --from=build-stage /app/bin/main /bin/main
CMD ["/bin/main"]