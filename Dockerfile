FROM golang:alpine
ENV GOOS=linux \
    GO111MODULE=on\
    GOPATH=""\
    CGO_ENABLE=0 \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn,direct"
RUN apk add build-base
RUN mkdir /consignment-service
COPY  ./src/consignment-service /consignment-service
RUN mkdir /vess-service
COPY  ./src/vess-service /vess-service

WORKDIR /consignment-service
RUN  go build -o consignment-service
WORKDIR ..
RUN mkdir /app
RUN cp /consignment-service/consignment-service  /app/consignment-service
WORKDIR /app
CMD ["./consignment-service"]