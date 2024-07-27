# Use an official Golang runtime as a base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /go/src/app

# Copy the entire project directory into the container's working directory
COPY . .

# Build the Go executable
RUN go build -o goshell main.go



# Command to run the executable
CMD ["./goshell"]
