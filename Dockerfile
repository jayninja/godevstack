# Part 1
FROM golang:latest AS builder

# Set working directory to /app (simple and direct)
WORKDIR /app

# Copy go.mod and go.sum from your source directory
COPY source/go.mod source/go.sum ./

# Download dependencies
RUN go mod download

# Copy all code from your source directory
COPY source/. .

# Build the application in /app
RUN go build -o main .


# Part 2
FROM golang:latest
WORKDIR /app
COPY --from=builder /app/main ./
EXPOSE 8080
CMD ["./main"]
