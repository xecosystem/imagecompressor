# Use the official Go image as the base image
FROM golang:1.16 AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Download and cache Go dependencies
RUN go mod download

# Copy the rest of the application code to the container
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o app

# Use a minimal base image for the final image
FROM scratch

# Copy the built binary from the previous stage
COPY --from=build /app/app /app

# Set the command to run when the container starts
CMD ["/app"]