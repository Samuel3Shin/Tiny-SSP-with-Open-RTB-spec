FROM golang:1.18

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY cmd/ssp/main.go .
COPY pkg/ssp/ pkg/ssp/
COPY pkg/common/ pkg/common/

RUN go build -o main

CMD ["./main"]
