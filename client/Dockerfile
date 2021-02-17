FROM golang:latest

# Set directory inside the container
WORKDIR $GOPATH/src/github.com/peter-crist/grpc-health-checks/client

# Copy everything from the current directory to the workdir in the container
COPY . .

# Download dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# Run the executable
CMD ["client"]

