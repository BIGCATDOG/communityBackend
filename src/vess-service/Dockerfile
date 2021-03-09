FROM golang:alpine
ENV GOOS=linux \
    GO111MODULE=on\
    GOPATH=""\
    CGO_ENABLE=0 \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn,direct"
RUN apk add build-base
COPY  . .
RUN  go build -o vess-service
RUN mkdir /app
RUN cp vess-service /app/vess-service
WORKDIR /app
CMD ["./vess-service"]