# Build Stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Download dependencies
COPY go.mod ./
# Missing go.sum locally due to environment issues, so we tidy inside container
RUN go mod tidy
RUN go mod download

# Build the application
COPY . .
RUN go build -o main .

# Run Stage
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
