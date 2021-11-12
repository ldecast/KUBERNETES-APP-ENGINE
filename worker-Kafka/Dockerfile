FROM golang:1.17

WORKDIR /app

ADD kafka ./kafka

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /go-kafka-worker

EXPOSE 8080


CMD [ "/go-kafka-worker" ]


