# Build stage
FROM golang:1.18 AS builder

# Set the working directory inside the builder container
WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Statically compile the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o my-echo-server ./cmd

# Final stage - minimal scratch image
FROM scratch

# Set the working directory in the scratch container
WORKDIR /root/

# Copy the statically compiled binary from the build stage
COPY --from=builder /app/my-echo-server .
COPY --from=builder /app/templates/.* templates/
COPY --from=builder /app/api/.* . api/

# Expose the port the server listens on
EXPOSE 8080

# Command to run the application
CMD ["./my-echo-server"]
