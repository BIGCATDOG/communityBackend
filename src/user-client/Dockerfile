FROM golang:alpine
ENV GOOS=linux \
    GO111MODULE=on\
    GOPATH=""\
    CGO_ENABLE=0 \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn,direct"
RUN apk add build-base
#RUN mkdir /user-client
#COPY  . /user-service
#
#WORKDIR /user-service
#RUN  go build -o user-service
#WORKDIR ..
COPY . .
RUN go build -o user-client
RUN mkdir /app
RUN cp user-client  /app/user-client
WORKDIR /app
CMD ["./user-client"]