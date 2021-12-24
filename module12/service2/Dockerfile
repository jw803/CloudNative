FROM golang:1.16-alpine as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o dist ./main.go

FROM alpine
COPY --from=builder /app/dist app
EXPOSE 80
ENTRYPOINT  ["./app"]