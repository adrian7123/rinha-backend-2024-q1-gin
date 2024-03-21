FROM golang:alpine AS builder

RUN apk update && apk add --no-cache 'git=~2'

ENV GO111MODULE=on
WORKDIR $GOPATH/src/packages/goginapp/
COPY . .

RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/main .

FROM alpine:3

WORKDIR /

COPY --from=builder /go/main /go/main

ENV PORT 8000
ENV GIN_MODE release

WORKDIR /go
COPY .env .

# Run the Go Gin binary.
ENTRYPOINT ["/go/main"]