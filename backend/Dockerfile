FROM golang:1.23-alpine

WORKDIR /app

# Install air for hot reload (allow Go toolchain auto-download)
ENV GOTOOLCHAIN=auto
RUN go install github.com/air-verse/air@latest

# Copy go mod files
COPY go.mod go.sum* ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

EXPOSE 8080

# Air will be run via docker-compose command
CMD ["air", "-c", ".air.toml"]
