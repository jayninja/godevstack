# Part 1
FROM golang:latest AS builder

WORKDIR /app
COPY source/go.mod source/go.sum ./
RUN go mod download
COPY source/. .
RUN go build -o main .


# Part 2
FROM golang:latest
WORKDIR /app
COPY --from=builder /app/main ./
EXPOSE 8080
CMD ["./main"]
