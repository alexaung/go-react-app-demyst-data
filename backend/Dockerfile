# Use the official Golang base image
FROM golang:1.17

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to the container's working directory
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code to the container's working directory
COPY . .

# Build the Go app
RUN go build -o main .

# Expose the port the app will run on
EXPOSE 8080

# Run the binary
CMD ["./main"]
