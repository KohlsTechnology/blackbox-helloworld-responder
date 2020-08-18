FROM gcr.io/distroless/static:nonroot

ARG HTTP_PORT=8080
ARG TCP_PORT=8081

ENV HELLO_WORLD_HTTP_PORT ${HTTP_PORT}
ENV HELLO_WORLD_TCP_PORT ${TCP_PORT}

COPY ./blackbox-helloworld-responder /blackbox-helloworld-responder

EXPOSE ${HTTP_PORT}/tcp
EXPOSE ${TCP_PORT}/tcp

ENTRYPOINT ["/blackbox-helloworld-responder"]
