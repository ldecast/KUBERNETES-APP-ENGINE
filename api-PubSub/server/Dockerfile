FROM golang:1.17

WORKDIR /app

ADD /games ./
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY key.json ./

RUN go build -o /go-pubsub-server

EXPOSE 9000

CMD [ "/go-pubsub-server" ]



