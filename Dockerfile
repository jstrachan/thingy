FROM golang:alpine AS build-env
WORKDIR /usr/local/go/src/../../../tmp/demo/thingy
COPY . /usr/local/go/src/../../../tmp/demo/thingy
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
RUN go get ./...
RUN go build -o build/thingy ./thingy


FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=build-env /usr/local/go/src/../../../tmp/demo/thingy/build/thingy /bin/thingy
CMD ["thingy", "up"]
