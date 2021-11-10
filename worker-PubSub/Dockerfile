FROM golang:1.17

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY key.json ./

RUN go build -o /go-pubsub-worker

EXPOSE 8080


CMD [ "/go-pubsub-worker" ]



