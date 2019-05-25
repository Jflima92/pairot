
FROM golang:1.12

# Set the Current Working Directory inside the container
WORKDIR /usr/src/app

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

RUN go get -d -v ./...

# Install the package
RUN go build -o builds/pairot ./cmd

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["./builds/pairot"]