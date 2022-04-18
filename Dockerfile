FROM golang:1.18 AS builder
WORKDIR /go/src/github.com/irvinlim/signalbin

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build \
    -a -v \
    -o bin/signalbin \
    ./...

FROM scratch

COPY --from=builder \
    /go/src/github.com/irvinlim/signalbin/bin/signalbin \
    /usr/local/bin/

ENTRYPOINT [ "/usr/local/bin/signalbin" ]
