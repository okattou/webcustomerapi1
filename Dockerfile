
FROM golang:1.19.5-alpine3.17

LABEL maintainer="BOTBIO <botbio.fr>"

RUN apk add --no-cache ca-certificates

USER root

ENV DBHOST 192.168.1.21
ENV HOST 172.17.0.2
ENV DBPORT 9006
ENV DBUSER postgres
ENV DBPASS postgres
ENV DBNAME postgres

COPY . /go/src/github.com/okattou/webcustomerapi1

#COPY ./launcher.sh \
#     /app/

EXPOSE 8081
WORKDIR /go/src/github.com/okattou/webcustomerapi1

CMD ["./run.sh"]

