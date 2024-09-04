FROM golang:1.23 AS build-stage
WORKDIR /app
ADD . .
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
RUN go install github.com/a-h/templ/cmd/templ@latest
RUN make build

FROM gcr.io/distroless/base-debian12 AS run-stage
COPY --from=build-stage /app/bin/main /bin/main
CMD ["/bin/main"]
