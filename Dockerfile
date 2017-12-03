FROM scratch

EXPOSE 8080

ADD terraform-server /
RUN chmod +x /terraform-server
CMD ["/terraform-server"]

ENTRYPOINT ["/terraform-server","--scheme=http", "--port=8080", "--host=0.0.0.0"]