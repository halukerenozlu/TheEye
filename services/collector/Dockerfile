# Build stage
FROM golang:1.25-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/collector ./cmd/collector

# Runtime stage
FROM alpine:3.20
WORKDIR /app
COPY --from=build /out/collector /app/collector
CMD ["/app/collector"]
