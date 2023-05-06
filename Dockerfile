# Build the executable under cmd/main.go
FROM golang:1.20.3-alpine3.17 as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o cmd/api main.go

# Use apline base image and copy the executable from builder
FROM alpine:3.17
RUN apk add --no-cache ca-certificates sed
# Copy the binary to the production image from the builder stage.
COPY  --from=builder /app/cmd/api /cmd/api
# Run the web service on container startup.
CMD ["/cmd/api"]
