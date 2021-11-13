FROM golang:1.17

WORKDIR /app

ADD games ./games
COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /go-rabbitmq-client

EXPOSE 10000

CMD [ "/go-rabbitmq-client" ]



