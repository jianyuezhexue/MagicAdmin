FROM golang:1.19
LABEL wangjie 490513975@qq.com
WORKDIR /go/src/magicAdmin

COPY . .

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env
RUN go mod tidy

EXPOSE 5240
ENTRYPOINT go run main.go