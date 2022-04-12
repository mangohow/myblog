# 该镜像用于编译web程序
FROM golang:alpine AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOPROXY="https://goproxy.cn,direct" \
    GO111MODULE=on

WORKDIR /build

COPY . .

RUN go mod download 
RUN go build -ldflags="-s -w" -o /app/blog_web 


# 该镜像用于运行web程序
FROM alpine

MAINTAINER mgh<xmg50120@hdu.edu.cn>

ENV TZ=Asia/Shanghai

WORKDIR /app

COPY --from=builder /app/blog_web /app/
COPY --from=builder /build/conf/conf.json /app/conf/

EXPOSE 8080

CMD ["./blog_web"]
