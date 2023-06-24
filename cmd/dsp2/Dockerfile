FROM golang:1.18

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY cmd/dsp2/main.go .
COPY pkg/dsp2/ pkg/dsp2/
COPY pkg/common/ pkg/common/

RUN go build -o main

CMD ["./main"]
