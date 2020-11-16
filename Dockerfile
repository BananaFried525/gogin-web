FROM golang:1.14

WORKDIR /go/src/app
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN go build

CMD ["./gogin-web"]