FROM golang:1.20rc2-alpine3.17
LABEL maintainer=wkl
ENV GO111MODULE=on \
    CGO_ENABLE=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn,direct"
# docker容器上新建opt目录
WORKDIR /opt
# 把dockerfile文件所在目录添加到 docker /opt下
ADD .  /opt
# 在容器内执行go 编译
RUN go build -o main ./main.go
# http服务监听端口，暴露给docker
EXPOSE 8080
# 运行go 程序的指令
CMD ["/opt/main"]