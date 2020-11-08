FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git 

RUN mkdir /GoSupervisor

WORKDIR /GoSupervisor

COPY . .

RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -installsuffix cgo -o /go/bin/GoSupervisor


FROM scratch

COPY --from=builder /go/bin/GoSupervisor /go/bin/GoSupervisor

ENV RABBITMQ_HOST=rabbitmq
ENV POSTGRES_HOST=postgres

ENTRYPOINT ["/go/bin/GoSupervisor"]

EXPOSE 8050