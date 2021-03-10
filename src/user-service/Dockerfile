FROM golang:alpine
ENV GOOS=linux \
    GO111MODULE=on\
    GOPATH=""\
    CGO_ENABLE=0 \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn,direct"
RUN apk add build-base
RUN mkdir /user-service
COPY  . /user-service

WORKDIR /user-service
RUN  go build -o user-service
WORKDIR ..
RUN mkdir /app
RUN cp /user-service/user-service  /app/user-service
WORKDIR /app
CMD ["./user-service"]