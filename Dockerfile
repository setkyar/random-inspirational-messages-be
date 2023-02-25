# Use the official Golang image as a parent image
FROM golang:1.17

# Set the working directory inside the container
WORKDIR /app

# Copy the necessary files to the container
COPY . .

# Build the Go app inside the container
RUN go build -o app

# Expose the port used by the app
EXPOSE 8080

# Run the app when the container starts
CMD ["./app"]