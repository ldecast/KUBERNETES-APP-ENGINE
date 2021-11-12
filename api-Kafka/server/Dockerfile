FROM golang:1.17

WORKDIR /app

ADD games ./games
ADD kafka ./kafka
COPY go.mod ./
COPY go.sum ./
RUN go mod download


COPY *.go ./

RUN go build -o /go-kafka-server

EXPOSE 9000

CMD [ "/go-kafka-server" ]


