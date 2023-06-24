FROM golang:1.18

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY cmd/logserver/main.go .
COPY pkg/logserver/ pkg/logserver/
COPY pkg/common/ pkg/common/

RUN go build -o main

CMD ["./main"]
