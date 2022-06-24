FROM golang:1.18-alpine

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

CMD ["go", "run", "./cmd/gophermart/main.go"]