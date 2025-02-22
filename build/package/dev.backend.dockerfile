FROM golang:latest

# Set the working directory in the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the rest of the application source code into the container
COPY . .

EXPOSE 8080

# Build and run the Go application
ENTRYPOINT ["go", "run", "./cmd/soarpipeline/soarpipeline.go"]