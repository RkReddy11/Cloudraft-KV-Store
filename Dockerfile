FROM golang:latest AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app .
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 8000
CMD ["./app"]

