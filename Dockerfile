FROM golang:alpine

ENV GIN_MODE=debug
ENV PORT=80

WORKDIR /chat_app/src

COPY src /chat_app/src

COPY assets /chat_app/assets
COPY templates /chat_app/templates
COPY js /chat_app/js

RUN go get ./...

RUN go build .

EXPOSE $PORT

ENTRYPOINT ["./chat_app"]