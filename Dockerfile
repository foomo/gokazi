FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN adduser -D -u 1001 -g 1001 gokazi

COPY gokazi /usr/bin/

USER gokazi
WORKDIR /home/gokazi

ENTRYPOINT ["gokazi"]
