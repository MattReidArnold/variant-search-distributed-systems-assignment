FROM golang:1.15

WORKDIR /app

COPY . .

RUN go get

CMD [ "go", "run", "main.go" ]
