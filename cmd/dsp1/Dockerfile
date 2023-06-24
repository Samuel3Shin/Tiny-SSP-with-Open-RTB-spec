FROM golang:1.18

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY cmd/dsp1/main.go .
COPY pkg/dsp1/ pkg/dsp1/
COPY pkg/common/ pkg/common/

RUN go build -o main

CMD ["./main"]
