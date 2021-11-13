FROM golang:1.17

WORKDIR /app


COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /go-rabbitmq-worker

EXPOSE 8080


CMD [ "/go-rabbitmq-worker" ]


