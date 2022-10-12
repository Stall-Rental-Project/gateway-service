FROM golang:1.19.0 as builder

ENV GOOS linux
ENV CGO_ENABLED 0
ENV GIN_MODE=release
ENV PORT=8000

WORKDIR /app
# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod download
# Copy app files
COPY . .
# Build app
RUN go build -o app

FROM alpine:3.16.2 as production
ENV GIN_MODE=release
ENV PORT=8000
# Add certificates
RUN apk add --no-cache ca-certificates

RUN mkdir assets
RUN cd assets
RUN mkdir images

# Copy built binary from builder
COPY --from=builder app .
# Expose port
EXPOSE $PORT
# Exec built binary
CMD ./app