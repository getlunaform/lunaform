FROM alpine:latest

EXPOSE 8080

RUN adduser -S lunarform

ADD lunarform /
RUN chmod +x /lunarform && \
    chown lunarform /lunarform

USER lunarform

CMD ["/lunarform"]

ENTRYPOINT ["/lunarform","--scheme=http", "--port=8080", "--host=0.0.0.0"]
