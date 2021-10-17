FROM golang:1.17 AS builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .
# Build the application
RUN go build -o main .

# Copy binary from build to main folder
RUN cp /build/main .

# # Build a small image
FROM scratch

COPY --from=builder /dist/ /

EXPOSE ${SERVER_PORT}
ENTRYPOINT ["./main"]

# # Command to run