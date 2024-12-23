FROM golang:1.23.4 AS builder

WORKDIR /go/src/github.com/KohlsTechnology/blackbox-helloworld-responder
COPY . .
RUN make build

FROM scratch

ARG HTTP_PORT=8080
ARG TCP_PORT=8081

ENV HELLO_WORLD_HTTP_PORT ${HTTP_PORT}
ENV HELLO_WORLD_TCP_PORT ${TCP_PORT}

EXPOSE ${HTTP_PORT}/tcp
EXPOSE ${TCP_PORT}/tcp

COPY --from=builder /go/src/github.com/KohlsTechnology/blackbox-helloworld-responder/blackbox-helloworld-responder /blackbox-helloworld-responder

ENTRYPOINT ["/blackbox-helloworld-responder"]
