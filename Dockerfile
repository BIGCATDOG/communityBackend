FROM golang:alpine
RUN apk add build-base
ENV GOOS=linux \
    GO111MODULE=on\
    GOPATH=""\
    CGO_ENABLE=0 \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn,direct"

COPY  ./src/consignment-service .
RUN   go build -ldflags="-s -w" -installsuffix cgo -o consignment-service main.go
RUN mkdir /app
WORKDIR /app
ADD consignment-service /app/consignment-service
CMD ["./consignment-service"]