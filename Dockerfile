FROM golang:1.21.3-alpine
WORKDIR /app
COPY go.mod ./
RUN go mod download

COPY *.go ./
RUN go test ./...

RUN go build .

CMD ./atlanta