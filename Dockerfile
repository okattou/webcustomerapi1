FROM golang:1.19.5-alpine3.17

LABEL maintainer="BOTBIO <botbio.fr>"

RUN apk add --no-cache ca-certificates

USER nobody

COPY ./launcher.sh \
     ./api-server \
     /app/

EXPOSE 3000

CMD ["/app/launcher.sh"]
