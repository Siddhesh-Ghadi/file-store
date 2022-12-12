FROM golang:1.19

WORKDIR $GOPATH/src/github.com/Siddhesh-Ghadi/file-store

COPY . .

RUN go build ./cmd/server/server.go

EXPOSE 8080

ENTRYPOINT ["./server"]
