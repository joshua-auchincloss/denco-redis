ARG DENCO_PORT=8080
FROM golang:latest

WORKDIR /usr/src/app

COPY go.mod go.sum ./
COPY ./main.go ./
COPY ./vendor ./
RUN go mod download && go mod verify

RUN go build -v -o /usr/local/bin/app ./...
RUN go run /usr/local/go/src/crypto/tls/generate_cert.go --ecdsa-curve P256 --host localhost
CMD ["app"]