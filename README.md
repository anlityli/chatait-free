# ChatAIT

<div align=center>
<img src="https://github.com/anlityli/chatait-free/blob/main/doc/logo.png?raw=true" width="200"/>
</div>

ChatAIT是用golang+vue开发的AI对话网站，对话模式模仿ChatGPT和Midjourney的形式。

正式版支持：
- [x] GPT4
- [x] GPT3.5
- [x] Midjourney
- [x] 充值和购买次数
- [x] openai接口循环调用
- [x] midjourney接口循环调用
- [x] midjourney接口多进程
- [x] midjourney U V Zoom Vary(strong) Vary(Subtle) ⬅️ ➡️ ⬆️ ⬇️ 等操作
- [x] midjourney 支持文生图和图生图 以及各种参数
- [x] midjourney 生成的图片可选是否本地保存
- [x] 管理员及角色管理
- [x] 会员管理
- [x] 会员等级
- [x] 对接微免签
- [x] midjourney提问支持百度翻译
- [x] midjourney图片本地保存缩略图，以便前台快速打开
- [ ] 敏感词过滤
- [ ] 站内公告文章系统
- [ ] 百度文心一言
- [ ] 更多实用功能

免费版支持：
- [x] GPT4
- [x] GPT3.5
- [x] 充值和购买次数
- [x] openai接口循环调用
- [x] 管理员及角色管理
- [x] 会员管理
- [x] 会员等级
- [ ] 更多实用功能

## 演示地址
[ChaiAIT官网](https://chat.chatait.top/)

## 截图
<div align=center>
<img src="https://github.com/anlityli/chatait-free/blob/main/doc/screenshot_1.jpg?raw=true" width="800"/>
<img src="https://github.com/anlityli/chatait-free/blob/main/doc/screenshot_2.png?raw=true" width="200"/>
<img src="https://github.com/anlityli/chatait-free/blob/main/doc/screenshot_3.png?raw=true" width="200"/>
</div>

## 视频部署教程
[https://www.bilibili.com/video/BV1gP411Y7oj](https://www.bilibili.com/video/BV1gP411Y7oj)

## Docker部署
[Docker部署](https://github.com/anlityli/chatait-free/blob/main/docker/README.md)

## 常规部署

### 解析域名
需要将以下域名解析到你服务器的IP，假设你的服务器主域名为```chatait.demo```
```
www.chatait.demo                # 网站前台域名（会员访问的域名）
backend.chatait.demo            # 网站后台域名（管理员访问的域名）
chat-frontend-api.chatait.demo  # 前台接口服务端域名（前台前端请求的接口域名）
chat-backend-api.chatait.demo   # 后台接口服务端域名（后台前端请求的接口域名）
```

> 注：如果部署免费版 ```chat-frontend-api``` 和 ```chat-backend-api``` 两个A记录必须严格按照这个内容解析。

### 购买服务器
最好购买国外服务器，因为国外服务器可以直接访问```openai```和```midjourney```网站，如果部署在国内服务器，部署完成后需要配置代理服务器。   
购买完成后安装linux系统，推荐centos7 或者 AlmaLinux 等RedHeat系统，我下面的步骤也都基于这些系统，有能力的可以采用其他系统部署。

下列服务器仅供参考，因为CN2 GIA 和 9929线路对国内访问速度友好。   
便宜服务器：[https://vmiss.com/](https://app.vmiss.com/aff.php?aff=409)   
8折优惠码：[20%off ](https://app.vmiss.com/aff.php?aff=409)

线路 | 内存 | CPU | SSD | 流量 | 带宽 | 价格 | 价格
---|---|---|---|---|---|---|---
美国 CN2 GIA |1G | 1核 | 10G | 400G/月 | 200M | 3.5加元/月 | [连接](https://app.vmiss.com/aff.php?aff=409&pid=1)
美国 CN2 GIA |1G | 1核 | 15G | 800G/月 | 200M | 5.6加元/月 | [连接](https://app.vmiss.com/aff.php?aff=409&pid=3)
美国 CN2 GIA |2G | 1核 | 20G | 1.2T/月 | 300M | 11加元/月 | [连接](https://app.vmiss.com/aff.php?aff=409&pid=4)
美国 CN2 GIA |2G | 2核 | 40G | 2T/月 | 500M | 21加元/月 | [连接](https://app.vmiss.com/aff.php?aff=409&pid=5)
美国 9929 |1G | 1核 | 10G | 500G/月 | 200M | 4加元/月 | [连接](https://app.vmiss.com/aff.php?aff=409&pid=57)
美国 9929 |1G | 1核 | 15G | 1T/月 | 200M | 6.8加元/月 | [连接](https://app.vmiss.com/aff.php?aff=409&pid=58)
美国 9929 |2G | 1核 | 20G | 1.5T/月 | 300M | 12.8加元/月 | [连接](https://app.vmiss.com/aff.php?aff=409&pid=59)
美国 9929 |2G | 2核 | 40G | 2.5T/月 | 500M | 24加元/月 | [连接](https://app.vmiss.com/aff.php?aff=409&pid=60)

> 注: 购买服务器后，设置的root密码一定要够复杂，尽量大于等于16个字符并包含英文大小写和数字，防止被暴力破解。

### 本地电脑安装ssh客户端工具
本地电脑安装ssh客户端工具，用于连接服务器   
windows: xshell, winscp, putty等   
mac: nuoshell, termius等   
这些工具随意，只要能连上服务器哪个都行

### 关闭selinux防火墙   
修改```/etc/selinux/config```文件 将 ```SELINUX=enforcing``` 改为 ```SELINUX=disabled```   
重启服务器
```
shutdown -r now
```

### 关闭firewall防火墙
执行命令
```
systemctl disable firewalld
```

### 新建chatait用户
新建一个非root用户，用该用户来启动服务端程序以保证网站的安全。执行以下命令
```
useradd chatait
passwd chatait
```
执行完这两条命令后，要求为```chatait```用户设置密码，输入过程没有回显，请一定注意，也是输入上了。输入完成后回车即可。然后接着输入第二次，之后回车。
创建完账户后，再执行以下命令来创建一个程序在服务器上的所在目录，执行：
```
mkdir /home/wwwroot/chatait
```
> 注: 设置的chatait密码一定要够复杂，尽量大于等于16个字符并包含英文大小写和数字，防止被暴力破解。

### 安装lnmp一键安装包
官网: [https://lnmp.org/](https://lnmp.org/)，依次执行下列命令：
```
cd ~ 
wget http://soft.lnmp.com/lnmp/lnmp2.0.tar.gz -O lnmp2.0.tar.gz && tar zxf lnmp2.0.tar.gz && cd lnmp2.0 && ./install.sh lnmp 
```
执行命令后，可能会出现一些报错提示，不要怕，提示缺少什么就安装什么，例如，缺少 ```wget```就安装，命令如下：
```
yum install wget
```
命令执行完成后会开始下载，下载完成后出现提示，依次按照需求选择要安装的MySQL版本及密码，以及PHP版本。 MySQL选择5.7或者8.0都可，然后输入MySQL的root密码，最后选择PHP版本，默认即可，因为本程序不需要PHP。 安装过程根据服务器和网络环境，持续时间大概在1小时到3小时不等。安装完成后会出现completed! enjoy it.字样。
> 注: 设置的数据库root密码一定要够复杂，尽量大于等于16个字符并包含英文大小写和数字，防止被暴力破解。

### 导入mysql数据库及初始数据
本地浏览器打开网址：```http://xxx.xxx.xxx.xxx/phpmyadmin``` 其中的xxx.xxx.xxx.xxx为服务器的ip地址，打开页面后```左侧菜单New```   

数据库名称: ```chatait_db```, 编码格式: ```utf8mb4_general_ci```

创建完成后，点击左侧菜单 ```chatait_db``` ， 点击顶部菜单```SQL```，把下载的程序包内 ```db```目录下的```chatait_db.sql```文件的内容粘贴到浏览器页面输入框内。点击```Go```按钮执行初始化数据库命令。

### 上传ChatAIT程序
下面利用winscp、lrzsz命令、scp命令，这些其一上传至服务器就好，如果是winscp直接拖拽到对应目录即可，如果是lrzsz命令的话，输入```rz```命令弹出提示框上传。      
将下载的程序包内 ```frontendServer``` ```backendServer``` ```frontendVue``` ```backendVue``` ```tmp``` ```files``` 这几个目录全部上传至服务器的```/home/wwwroot/chat/``` 目录下。   

### 配置文件
打开本地程序包内```config```目录下的```config.toml```文件，把最下面的数据库配置部分按照你的服务器实际用户名密码来修改，如果其他步骤都和我的流程一致，则其他部分无需修改。   
修改完成后，将```config```目录上传至```/home/wwwroot/chat/```下。

### 配置nginx
nginx是会员能访问到网站的关键步骤，一定要按照你的真实域名修改。打开本地程序包内的```nginx```目录下的```chatait.conf```文件，找到所有的```chatait.demo```替换为你的主域名，如果按照我的流程部署，其他位置无需修改。   
修改完成后，将```chatait.conf```文件上传至```/usr/local/nginx/conf/vhost/```下。

> 注：如果你的网站想要启用ssl即https，需要注释掉监听的80端口，取消443端口监听前面的注释。并把ssl相关的注释打开，配置好你的ssl证书的key和crt文件路径。

配置完成后，执行重启nginx命令
```
lnmp nginx restart
```

### 修改程序目录的权限
把程序目录的权限全部改成```chatait```的权限，防止root启动程序带来安全隐患。执行以下命令：
```
chown -R chatait:chatait /home/wwwroot/chatait
chmod -R 755 /home/wwwroot/chatait
```

### 检查网站部署情况
用```chatait```用户登录ssh终端，执行以下命令来检查前端部署情况：
```
/home/wwwroot/chatait/frontendServer/chatait-frontend-server --gf.gcfg.path=/home/wwwroot/chatait/config
```
执行该命令后，打开浏览器，打开网址 ```http://www.chatait.demo``` ，如果网站能正常访问，点击登录按钮有反应(提示用户名和密码错误也是反应)，则证明前台部署成功，按```ctrl```+```C``` 按键打断命令。

接下来执行以下命令来检查后台部署情况：
```
/home/wwwroot/chatait/backendServer/chatait-backend-server --gf.gcfg.path=/home/wwwroot/chatait/config
```
执行该命令后，打开浏览器，打开网址 ```http://backend.chatait.demo``` ，如果网站能正常访问，点击登录按钮有反应(提示用户名和密码错误也是反应)，则证明后台部署成功，按```ctrl```+```C``` 按键打断命令。

> 注：默认管理员账户: admin, 密码: admin111

### 配置服务项并设置开机启动
上面的步骤执行完成后，如果ssh客户端一旦退出，程序就会中断，网站服务端接口将无法访问。因为我们需要将前台服务端和后台服务端加入到系统项里，并配置开机启动。   
首先，以```root```用户登录ssh客户端，将程序包内的```system```目录内的 ```chataitBackend.service``` 和 ```chataitFrontend.service``` 两个文件上传至服务器 ```/etc/systemd/system``` 目录下，然后依次执行以下命令：
```
systemctl daemon-reload
systemctl enable chataitBackend.service
systemctl enable chataitFrontend.service
systemctl start chataitBackend.service
systemctl start chataitFrontend.service
```
上述命令执行完成后，再次浏览器打开前台和后台网址，看看是否能正常访问，并且点击登录按钮都有反应(提示用户名和密码错误也是反应)，则全部部署完毕。

## 注意事项
以上步骤没有涉及到防火墙等一些其他安全措施，还可以做进一步的安全调整，这个需要有一定服务器运维经验才能完成，所以，一旦有了一定的运维经验，也就不用非常严格的按照我的步骤来部署了。为了让新手能够以最快的速度和最简单的流程完成部署，所以我写的这些流程省去了一些安全设置。

## 捐赠
<div align=center>
<img src="https://github.com/anlityli/chatait-free/blob/main/doc/donate.jpg?raw=true" width="375"/>
</div>

## 联系QQ购买正式版或定制开发
邮箱: leo@leocode.net   
QQ: 88511136

## 框架
golang: [Goframe](https://github.com/gogf/gf)   
Vue: [TDesign](https://tdesign.tencent.com/)



