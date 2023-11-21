# Start from the latest golang base image for the builder stage
FROM golang:latest as builder

# Add Maintainer Info
LABEL maintainer="DKZhang <527656264@qq.com>"

# Set the Current Working Directory inside the builder container
WORKDIR /app

# Install git.
RUN apt-get update && apt-get install -y git

# Clone the repository
RUN git clone https://github.com/dkzhang/mrm.git . #20231121-1103

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

# Build the Go app. Note that CGO is disabled to create a static binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the pre-built binary from the previous stage
COPY --from=builder /app/main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]