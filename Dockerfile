FROM golang:alpine

ENV GIN_MODE=release
ENV PORT=80

WORKDIR /chat_app/src

COPY src /chat_app/src

COPY assets /chat_app/assets
COPY templates /chat_app/templates

# RUN apk update && apk add --no-cache git
RUN go get ./...

RUN go build .

EXPOSE $PORT

ENTRYPOINT ["./chat_app"]