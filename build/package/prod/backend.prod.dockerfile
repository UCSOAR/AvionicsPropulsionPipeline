# Use the latest Go image if root access is available and version flexibility is acceptable.
# Use a specific Go version for consistency and when rootless Docker environments require it.
FROM golang:1.23 AS build

# Set the working directory in the container
WORKDIR /app

# Copy the Go module files. Make sure go.sum exists. If not, or if there's an error, run `go mod tidy` locally.
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code into the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/soarpipeline ./cmd/soarpipeline

# Create a minimal image to run the Go application
FROM alpine:latest

# Install necessary CA certificates
RUN apk --no-cache add ca-certificates

# Set the working directory in the container
WORKDIR /root/

# Copy the .env.toml file
COPY .env.toml .

# Copy the built Go application from the build stage
COPY --from=build /app/soarpipeline .

# Ensure the binary has executable permissions
RUN chmod +x ./soarpipeline

# Expose port 8080
EXPOSE 8080

# Run the Go application
ENTRYPOINT ["./soarpipeline"]