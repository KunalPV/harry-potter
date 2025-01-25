# Use Go's official image
FROM golang:1.20

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Install dependencies
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the application
RUN go build -o main .

# Expose the app on port 8080
EXPOSE 8080

# Run the app
CMD [ "./main" ]