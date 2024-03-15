FROM node:16.15-alpine
WORKDIR /app
COPY ./front-end ./
RUN npm install
RUN npm run build

FROM golang:1.21.3-alpine
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY .env ./
ADD . .

COPY --from=0 /app/build ./front-end/build

RUN go build .

CMD ./atlanta