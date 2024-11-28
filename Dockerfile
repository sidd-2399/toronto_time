# Use the official Go image to build the application
FROM golang:1.21-alpine as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go mod and sum files
COPY go.mod ./

# Download Go dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o main .

# Use a smaller image to run the application
FROM alpine:latest

# Install necessary dependencies (like certificates for secure connections)
RUN apk --no-cache add ca-certificates

# Set the working directory inside the container
WORKDIR /root/

# Copy the compiled Go binary from the builder image
COPY --from=builder /app/main .

# Expose the port the app runs on
EXPOSE 8080

# Run the Go binary when the container starts
CMD ["./main"]
