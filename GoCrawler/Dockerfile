FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git && apk add curl

RUN mkdir /goCrawler

WORKDIR /goCrawler

COPY . .

RUN go mod download

ENV RABBITMQ_HOST=rabbitmq
ENV POSTGRES_HOST=postgres

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -installsuffix cgo -o /go/bin/goCrawler

ENTRYPOINT ["/go/bin/goCrawler"]

EXPOSE 8080
