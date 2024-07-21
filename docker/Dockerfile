FROM golang:1.22-alpine

## Build dependencies
RUN apk add --no-cache alpine-sdk

WORKDIR /build

## Build dependencies
RUN apk add --no-cache alpine-sdk

## Application dependencies
ADD go.mod go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY internal ./internall

RUN go build -C cmd

CMD [ "./cmd/cmd" ]
