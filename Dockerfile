FROM alpine:latest

EXPOSE 8080

RUN adduser -S lunaform

ADD lunaform-server /lunaform/lunaform-server
RUN chmod +x /lunaform/lunaform-server && \
    chown lunaform /lunaform/lunaform-server

USER lunaform

ENTRYPOINT ["/lunaform/lunaform-server","--scheme=http", "--port=8080", "--host=0.0.0.0"]
