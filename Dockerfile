FROM golang:1.23 AS builder
WORKDIR /application
COPY  go.* .
RUN go mod download # download dependencies
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o app ./cmd/app

FROM alpine:3.19.1
RUN apk --no-cache add ca-certificates
WORKDIR /application
COPY --from=builder /application/app .
CMD ["/application/app"]
