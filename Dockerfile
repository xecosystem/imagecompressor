# Use the official Go image as the base image
FROM golang:1.18 AS build

# Install ca-certificates
RUN apt-get update && apt-get install -y ca-certificates

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Download and cache Go dependencies
RUN go mod download

# Copy the rest of the application code to the container
COPY . .

# Download msgp
RUN go mod download github.com/tinylib/msgp

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o app

# Use a minimal base image for the final image
FROM scratch

# Copy the ca-certificate.crt from the build stage
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy the built binary from the previous stage
COPY --from=build /app/app /app

# Set the command to run when the container starts
CMD ["/app"]
