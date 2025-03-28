FROM golang:1.23-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./

RUN go env -w GOPROXY=https://goproxy.cn,direct

RUN go mod download

COPY . .

RUN go build -o oss-srv app/oss/main.go

FROM alpine:3.19

WORKDIR /build

# 从构建阶段复制二进制文件
COPY --from=builder /build/oss-srv ./oss-srv

# 复制配置文件（可选）
COPY app/oss/oss.yaml ./

EXPOSE 50053

CMD ["./oss-srv"]
