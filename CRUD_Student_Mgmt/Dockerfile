FROM golang:latest

WORKDIR /myapp

COPY . .

RUN go get -u google.golang.org/grpc && \
    go get -u github.com/golang/protobuf/protoc-gen-go

RUN apt-get update && \
    apt-get install -y zip && \
    mkdir /opt/protoc && \
    cd /opt/protoc && \
    wget https://github.com/protocolbuffers/protobuf/releases/download/v3.7.0/protoc-3.7.0-linux-x86_64.zip && \
    unzip protoc-3.7.0-linux-x86_64.zip && \
    rm protoc-3.7.0-linux-x86_64.zip

ENV PATH=$PATH:$GOPATH/bin:/opt/protoc/bin

RUN export PATH="/usr/bin/protoc:$PATH"

CMD ["go", "run", "./Server/main.go"]

EXPOSE 50051
