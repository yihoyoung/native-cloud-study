FROM golang AS builder

RUN go build -o /usr/bin/go-server

FROM alpine

COPY --form=builder /usr/bin/go-server /usr/bin

ENTRYPOINT ["/usr/bin/go-server"]