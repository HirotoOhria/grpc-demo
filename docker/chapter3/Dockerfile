FROM golang:1.17

WORKDIR /go/src

RUN apt update && \
    apt install -y protobuf-compiler

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

CMD ["go", "run", "server/server.go"]
