FROM scratch
ADD ca-certificates.crt /etc/ssl/certs/
ADD terraform-server /
CMD ["/terraform-server","--scheme=http"]