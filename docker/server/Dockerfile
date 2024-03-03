FROM --platform=$BUILDPLATFORM golang:1.18 as builder

ARG TARGETARCH
ARG GOPROXY=https://goproxy.cn,direct

WORKDIR /app
RUN mkdir /app/backendServer
RUN mkdir /app/frontendServer

COPY ./backendServer /app/backendServer/
COPY ./frontendServer /app/frontendServer/

WORKDIR /app/backendServer
RUN GOOS=linux GOARCH=$TARGETARCH go build -a -o /app/backendServer/chatait-backend-server main.go
WORKDIR /app/frontendServer
RUN GOOS=linux GOARCH=$TARGETARCH go build -a -o /app/frontendServer/chatait-frontend-server /main.go

FROM alpine:3.18

WORKDIR /chatait
RUN mkdir /chatait/config
RUN mkdir /chatait/tmp/
RUN mkdir /chatait/files/
RUN mkdir /chatait/backendServer
RUN mkdir /chatait/backendServer/runtime
RUN mkdir /chatait/frontendServer
RUN mkdir /chatait/frontendServer/runtime

VOLUME /chatait/config
VOLUME /chatait/tmp/
VOLUME /chatait/files/

COPY --from=builder /app/backendServer/chatait-backend-server /chatait/backendServer
COPY --from=builder /app/frontendServer/chatait-frontend-server /chatait/frontendServer
COPY ./docker-run.sh /chatait

WORKDIR /chatait
CMD ["/bin/sh", "-c", "/chatait/docker-run.sh"]


# FROM alpine:3.18
# 
# WORKDIR /chatait
# RUN mkdir /chatait/config
# RUN mkdir /chatait/tmp/
# RUN mkdir /chatait/files/
# RUN mkdir /chatait/backendServer
# RUN mkdir /chatait/backendServer/runtime
# RUN mkdir /chatait/frontendServer
# RUN mkdir /chatait/frontendServer/runtime
# VOLUME /chatait/config
# VOLUME /chatait/tmp/
# VOLUME /chatait/files/
# COPY ./backendServer/chatait-backend-server /chatait/backendServer/
# COPY ./frontendServer/chatait-frontend-server /chatait/frontendServer/
# COPY ./docker-run.sh /chatait
# RUN chmod +x /chatait/backendServer/chatait-backend-server
# RUN chmod +x /chatait/frontendServer/chatait-frontend-server
# 
# WORKDIR /chatait
# CMD ["/bin/sh", "-c", "/chatait/docker-run.sh"]

