FROM golang:1.17-alpine

WORKDIR /app

COPY gateway/go.mod gateway/go.sum ./
RUN go mod download

COPY . .

RUN go build -o gateway ./gateway/cmd

EXPOSE 8080

CMD ["./gateway"]
