FROM golang:1.23 AS build-stage
WORKDIR /app
ADD . .
RUN make build

FROM gcr.io/distroless/base-debian12 AS run-stage
COPY --from=build-stage /app/bin/main /bin/main
CMD ["/bin/main"]
