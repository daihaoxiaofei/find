FROM alpine

WORKDIR /work

COPY find .
COPY config/config.yaml config/config.yaml

RUN chmod +x find

EXPOSE 8080
CMD ./find


# docker build -t daihaoxiaofei/find:0.0.1 .     # 打包


# 登录  docker login daihaoxiaofei
# 上传  docker push daihaoxiaofei/find:0.0.1