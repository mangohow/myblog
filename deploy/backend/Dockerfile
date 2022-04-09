# 该镜像用于编译web程序
FROM golang:alpine AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOPROXY="https://goproxy.cn,direct" \
    GO111MODULE=on

WORKDIR /build

COPY . .

RUN go mod download \
    go build -ldflags="-s -w" -o /app/blog_web ./main.go


# 该镜像用于运行web程序
FROM alpine

MAINTAINER mgh<xmg50120@hdu.edu.cn>

ENV TZ=Asia/Shanghai

WORKDIR /app

COPY --from=builder /app/blog_web /app/blog_web

EXPOSE 8080

CMD ["./blog_web"]