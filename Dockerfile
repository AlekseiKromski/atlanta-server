FROM golang:1.21.3-alpine
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY .env ./
COPY *.go ./
RUN go test ./...

RUN go build .

CMD ./atlanta