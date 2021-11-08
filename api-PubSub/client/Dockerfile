FROM golang:1.17

WORKDIR /app

ADD /games ./
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /go-pubsub-client

EXPOSE 10000

CMD [ "/go-pubsub-client" ]











