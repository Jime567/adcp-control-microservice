FROM alpine:3.18

RUN apk update && apk add bash && apk --no-cache add tzdata

ARG NAME

COPY ${NAME} /adcp-control

ENTRYPOINT ["/adcp-control"]