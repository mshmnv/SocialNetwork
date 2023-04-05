FROM golang:1.20.2-alpine

ENV GOPATH=/

COPY ./ ./

RUN go mod download

RUN go build ./cmd/server

ENTRYPOINT [ "./server" ]
