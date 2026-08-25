# ---------- Build stage ----------
FROM golang:1.26-alpine AS builder

WORKDIR /app

# Copy dependency file first
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy application source
COPY . .

# Build the Go binary
RUN go build -o app .


# ---------- Runtime stage ----------
FROM alpine:3.22

WORKDIR /app

# Copy only the compiled binary
COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]