FROM        golang:1.17.2-stretch as builder
RUN         mkdir -p /app
WORKDIR     /app
COPY        . .
RUN  go build -o /app

FROM bitnami/minideb:stretch
RUN mkdir -p /app
WORKDIR /app
COPY --from=builder /app /app

ENTRYPOINT  ["./app"]