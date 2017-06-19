FROM alpine:latest
MAINTAINER Charles Teinturier "teintu.c@gmail.com"

COPY webconfig  /opt/

EXPOSE 80:8080

ENTRYPOINT ["/opt/webconfig"]

COPY Dockerfile /
