FROM golang:1.14 AS builder

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.io,direct

# 移动到工作目录：/build
WORKDIR /build

# 复制项目中的 go.mod 和 go.sum文件并下载依赖信息
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o sp500_straddle main.go

###################
# 接下来创建一个小镜像
###################
FROM debian:stretch-slim

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.io,direct

COPY ./wait-for.sh /


# 从builder镜像中把/dist/app 拷贝到当前目录
COPY --from=builder /build/sp500_straddle /

RUN sed -i s@/deb.debian.org/@/mirrors.aliyun.com/@g /etc/apt/sources.list \
&& apt-get clean \
&& apt-get update;

RUN apt-get install -y \
		--no-install-recommends \
		netcat; \
        chmod 755 wait-for.sh

# CMD ["go", "run", "main.go"]