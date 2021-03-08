FROM golang:alpine
ENV GOOS=linux \
    GO111MODULE=on\
    GOPATH=""\
    CGO_ENABLE=0 \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn,direct"
RUN apk add build-base
COPY  ./src/consignment-service .
RUN  go build -o consignment-service
RUN mkdir /app
RUN cp consignment-service /app/consignment-service
WORKDIR /app
CMD ["./consignment-service"]