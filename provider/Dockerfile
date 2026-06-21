############################################################
# Dockerfile to build gRPC-Go-Polaris quickstart example provider

# First of all, you need to build the binary from the source code,
# use `CGO_ENABLED=0 go build -ldflags "-s -w" -o provider` to build the binary.

############################################################

FROM golang:alpine as builder

ENV WORKDIR /root/app

#ENV GOPROXY=https://goproxy.cn,https://goproxy.io,direct \
#  GO111MODULE=on 

ENV TZ=Asia/Shanghai

WORKDIR $WORKDIR
ADD ./polaris.yaml $WORKDIR
ADD ./bin $WORKDIR
RUN chmod +x $WORKDIR/bin

# ENTRYPOINT ./provider

CMD /root/app/bin
