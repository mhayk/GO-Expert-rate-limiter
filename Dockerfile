FROM golang:1.22.5-alpine

WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./
COPY go.sum ./

# Download all dependencies of the Go module
RUN go mod tidy
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN go build -o main ./cmd/server

# Execute the Go app
CMD ["./main"]