FROM golang:1.17-alpine

WORKDIR /usr/src/app

COPY micro1/go.mod micro1/go.sum ./
RUN go mod download

COPY . .

RUN go build -o grpc-server ./micro1/cmd

EXPOSE 50051

CMD ["./grpc-server"]