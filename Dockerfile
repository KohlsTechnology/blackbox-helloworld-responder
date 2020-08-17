FROM gcr.io/distroless/static:nonroot

COPY ./blackbox-helloworld-responder /blackbox-helloworld-responder

ENV HELLO_WORLD_HTTP_PORT=8080
ENV HELLO_WORLD_TCP_PORT=8081

EXPOSE 8080/tcp
EXPOSE 8081/tcp

ENTRYPOINT ["/blackbox-helloworld-responder"]
