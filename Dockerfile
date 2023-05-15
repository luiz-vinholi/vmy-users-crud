FROM golang:1.20

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -v -o main ./main.go

EXPOSE 8888

CMD ["./main"]