FROM --platform=linux/amd64 golang:1.20
MAINTAINER OD.Ice
WORKDIR /go/src/survey_backend
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...

# 安装 supervisor
RUN apt-get update && apt-get install -y supervisor

# 安装git
RUN apt-get update && apt-get install -y git

# 复制 supervisor 配置文件到容器中
COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf

# 运行 supervisor
CMD ["supervisord", "-c", "/etc/supervisor/conf.d/supervisord.conf"]