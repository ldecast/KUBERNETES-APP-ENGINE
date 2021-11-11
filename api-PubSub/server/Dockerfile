FROM golang:1.17

WORKDIR /app

ADD games ./games
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY key.json ./
COPY *.go ./

RUN go build -o /go-pubsub-server

EXPOSE 9000

CMD [ "/go-pubsub-server" ]


