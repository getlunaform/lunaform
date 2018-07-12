FROM alpine:latest

EXPOSE 8080

RUN adduser -S lunaform

ADD lunaform /
RUN chmod +x /lunaform && \
    chown lunaform /lunaform

USER lunaform

CMD ["/lunaform"]

ENTRYPOINT ["/lunaform","--scheme=http", "--port=8080", "--host=0.0.0.0"]
