# Docker部署说明

## 部署mysql
这一步自行部署mysql并导入初始化数据，可以参考[常规部署](https://github.com/anlityli/chatait-free/blob/main/README.md#导入mysql数据库及初始数据)的mysql部分。

## 拉取镜像
```
docker pull anlity/chataitserver:1.1.1
docker pull anlity/chataitvue:1.1.1
```

## 启动服务端docker
```
docker run -p 18001:18001 -p 18002:18002 -v /root/chatait-free/docker/server/config:/chatait/config -v /root/chatait-free/docker/server/tmp:/chatait/tmp/ -v /root/chatait-free/docker/server/files:/chatait/files/ -d --name chataitServerApp anlity/chataitserver:1.1.1
```
```/root/chatait-free/docker/server/config```，```/root/chatait-free/docker/server/tmp```，```/root/chatait-free/docker/server/files``` 这三个为本地宿主机目录，用来存放配置文件和midjourney生成的临时文件以及储存目录。   

注意把chatait-free/docker/server/config内的文件复制到宿主机对应的config目录内，根据实际情况配置，一般只需要配置mysql部分即可

## 启动前端docker
```
docker run -p 18003:18003 -p 18004:18004 -d --name chataitVueApp anlity/chataitvue:1.1.1
```

## 部署nginx
复制```chatait-free/docker/chatait.conf```内容到nginx配置文件中，可以参考[常规部署](https://github.com/anlityli/chatait-free/blob/main/README.md#配置nginx)的nginx部分。   
注意，修改域名为你的域名。