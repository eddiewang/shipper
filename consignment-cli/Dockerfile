FROM golang:latest as builder

WORKDIR /go/src/github.com/eddiewang/shipper/consignment-cli/consignment-cli

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep

RUN dep ensure

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM alpine:latest
# Security related package, good to have.
RUN apk --no-cache add ca-certificates

RUN mkdir -p /app
WORKDIR /app

# Here, instead of copying the binary from our host machine,
# we pull the binary from the container named `builder`, within
# this build context. This reaches into our previous image, finds
# the binary we built, and pulls it into this container. Amazing!
COPY --from=builder /go/src/github.com/eddiewang/shipper/consignment-cli/consignment-cli .

CMD ["./consignment-cli"]