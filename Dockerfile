FROM golang:1.18

RUN apt-get update && apt-get install -y \
    iproute2 inetutils-ping net-tools tcpdump netcat nmap \
    entr

WORKDIR /go/src

COPY go.mod /go/src/go.mod
COPY go.sum /go/src/go.sum

RUN go mod download

COPY internal /go/src/internal
COPY pkg /go/src/pkg
COPY main.go /go/src/main.go

STOPSIGNAL SIGKILL

ENTRYPOINT []
CMD []
