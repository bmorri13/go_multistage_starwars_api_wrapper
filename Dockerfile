# Start with a golang base image to build our application
FROM cgr.dev/chainguard/go:latest as build-env

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./
COPY go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 go build -o /go-server

# Start a new static cgr image to reduce the size of the final image
FROM cgr.dev/chainguard/static:latest

# Copy the Pre-built binary file from the previous stage
COPY --from=build-env /go-server /

# Expose port 5002 to the outside world
EXPOSE 5002

# Command to run the executable
ENTRYPOINT ["/go-server"]
