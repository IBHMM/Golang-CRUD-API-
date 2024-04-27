FROM golang:1.22.2

WORKDIR /app

COPY go.mod .
COPY main.go .

RUN  go build -o bin .

ENTRYPOINT [ "/app/bin" ]