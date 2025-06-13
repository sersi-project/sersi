FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy go files
COPY . .

# Build the Go application
RUN go build -o /bin/sersi

# Create a new stage for the final image
FROM node:20-alpine

# Copy the Go binary from builder stage
COPY --from=builder /bin/sersi /usr/local/bin/sersi

# Copy the test script
COPY test/integration-test.sh .

# Make the script executable
RUN chmod +x integration-test.sh

# Run the test script
ENTRYPOINT ["./integration-test.sh"]
