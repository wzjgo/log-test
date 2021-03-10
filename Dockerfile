FROM golang:1.15-alpine AS builder
RUN apk add --no-cache ca-certificates git

WORKDIR /
# restore dependencies
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o /app .

FROM alpine AS release
RUN apk add --no-cache ca-certificates
WORKDIR /logtest
COPY --from=builder /app ./app
ENTRYPOINT ["/logtest/app"]

