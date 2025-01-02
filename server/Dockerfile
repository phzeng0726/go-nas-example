FROM golang:alpine AS builder

WORKDIR /app

COPY . /app

# 移動檔案到 server 目錄
RUN mkdir ./server && mv cmd go.* ./server

# 建置二進位檔案
ENV CGO_ENABLED=1
RUN apk add --no-cache gcc musl-dev
RUN cd ./server && go get -u ./... && go mod tidy
RUN cd ./server && GOOS=linux go build -v -o ../main ./cmd


# 最終映像階段
FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache tzdata

COPY --from=builder /app/main /app/main

EXPOSE 8080

# 執行應用程式
CMD ["/app/main", "-c"]