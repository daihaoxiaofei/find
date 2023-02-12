# 打包依赖阶段使用golang作为基础镜像
FROM golang:1.18.8-alpine3.17 as builder

# 启用go module
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPATH=/GOPATH \
    GOPROXY="https://goproxy.cn,direct"

# 事先复制过来可以利用docker的缓存 只要mod文件不动就可以不重新下载依赖
COPY go.mod go.sum ./
RUN go mod download

# 编译
COPY . .
RUN go build -a -installsuffix cgo -o find main.go


# 运行阶段指定 scratch或alpine 作为基础镜像
FROM alpine:3.17

# 换源 安装 时区
RUN echo -e http://mirrors.ustc.edu.cn/alpine/v3.12/main/ > /etc/apk/repositories &&\
    apk add --no-cache  tzdata &&\
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

# 将上一个阶段文件夹下的编译文件复制进来 默认目录为/go
COPY --from=builder /go/find .

# 运行
CMD ./find
