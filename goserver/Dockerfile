ARG DENCOPORT=8080
FROM golang:latest

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
COPY server.go ./
RUN go mod download && go mod verify

RUN go build -v -o /usr/local/bin/app ./...

EXPOSE ${DENCOPORT}
CMD ["app"]