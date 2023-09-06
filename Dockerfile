# Use an official Go runtime as a parent image
FROM golang:1.21

# Set the working directory inside the container
WORKDIR /app

# Copy the local code to the container
COPY . .

# Build the Go application
RUN go build -o build cmd/server/main.go

# Expose a port for your Go application
EXPOSE 8085

# Define the command to run your application
CMD ["./build"]
