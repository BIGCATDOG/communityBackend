FROM golang:alpine
ENV GOOS=linux \
    GO111MODULE=on\
    GOPATH=""\
    CGO_ENABLE=0 \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn,direct"
RUN apk add build-base
RUN   mkdir client
COPY  /client  /client
RUN   mkdir -p /consignment-service/proto
COPY  consignment-service/proto /consignment-service/proto

WORKDIR  /client
RUN  go build -o client
WORKDIR ../
RUN mkdir /app
RUN cp /client/client /app/client
RUN cp /client/consignment.json /app/consignment.json
WORKDIR /app
CMD ["./client"]