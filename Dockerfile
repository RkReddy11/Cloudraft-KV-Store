FROM golang:latest AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

# Create a minimal runtime image
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 8000
CMD ["./app"]

