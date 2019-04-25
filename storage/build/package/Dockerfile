# builder image
FROM golang:1.12 as builder

WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download
RUN mkdir /build

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /build/aggregator ./cmd/storage

# minimal image
FROM alpine:latest

RUN apk update && \
   apk add ca-certificates && \
   update-ca-certificates && \
   rm -rf /var/cache/apk/*

COPY --from=builder /build /app
WORKDIR /app

EXPOSE 80

ENTRYPOINT ["/app/storage"]