FROM golang:1.16-alpine as builder
WORKDIR /app
COPY . .
RUN go build -o dist ./main.go

FROM alpine
COPY --from=builder /app/dist /app
COPY config.yaml /app
EXPOSE 9090
ENTRYPOINT  ["./app"]