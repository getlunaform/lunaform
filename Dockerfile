FROM alpine:latest

EXPOSE 8080

RUN adduser -S terraform-server

ADD terraform-server /
RUN chmod +x /terraform-server && \
    chown terraform-server /terraform-server

USER terraform-server

CMD ["/terraform-server"]

ENTRYPOINT ["/terraform-server","--scheme=http", "--port=8080", "--host=0.0.0.0"]