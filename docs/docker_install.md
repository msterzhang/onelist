## 1.拉取docker镜像
```
docker pull msterzhang/onelist:latest
```

## 2.新建一个用于保存配置相关文件的目录，比如：
```
/root/onelist/config
```

## 3.运行
```
docker run -d --name onelist -e PUID=0 -e PGID=0 -e TZ=Asia/Shanghai -p 5245:5245 -v /root/onelist/config:/config msterzhang/onelist:latest
```

## 4.修改配置
编辑`/root/onelist/config`目录下config.env
```

# 服务设置
# 注意要改为未被占用的端口
API_PORT=5245
FaviconicoUrl=https://wework.qpic.cn/wwpic/818353_fizV30xbQCGPQRP_1677394564/0
API_SECRET=8qQU6Uz0pgF8pYOK1huPUnk2nopM3DHi

# 网站名称
Title=onelist

# Env有两种模式，Debug及Release，主要用在数据库为mysql时候，需要注意修改Env环境和mysql密码对应
Env=Debug

# 管理员账户设置，用于初始化管理员账户
UserEmail=xxxx.@qq.com
UserPassword=xxxxx

# 下载刮削图片到本地
DownLoadImage=是
# 留空则表示使用本地缓存图片,否则使用https://image.tmdb.org
ImgUrl=
# 允许刮削alist中的视频文件类型
VideoTypes=.mp4,.mkv,.flv

# 数据库设置
DB_DRIVER=sqlite
DB_USER=root
DbName=onelist

# 如果上面DB_DRIVER类型为mysql，就需要正确填下以下参数
DB_PASSWORD_Debug=123456
DB_PASSWORD_Release=123456

# TheMovieDb Key
# 在https://www.themoviedb.org网站申请
KeyDb=22f10ca52f109158ac7fe064ebbcf697
```
## 5.重启容器
```
docker restart onelist
```
## 6.国内主机要修改Docker容器的hosts文件，才可以刮削资源：

### 进入运行中的Docker容器：
```
docker exec -it onelist /bin/bash
```
### 编辑hosts文件：
```
vi /etc/hosts
```
### 在文件的末尾添加所需的主机名和IP地址：

```
13.224.161.90 api.themoviedb.org
104.16.61.155 image.themoviedb.org
13.35.67.86 api.themoviedb.org
54.192.151.79 www.themoviedb.org
13.225.89.239 api.thetvdb.com
13.249.175.212 api.thetvdb.com
13.35.161.120 api.thetvdb.com
13.226.238.76 api.themoviedb.org
13.35.7.102 api.themoviedb.org
13.225.103.26 api.themoviedb.org
13.226.191.85 api.themoviedb.org
13.225.103.110 api.themoviedb.org
52.85.79.89 api.themoviedb.org
13.225.41.40 api.themoviedb.org
13.226.251.88 api.themoviedb.org
```
### 保存并退出文件。
```
按Esc键，进入vim命令模式
输入
:wq
回车退出
```
### 退出容器：
```
exit
```
现在，您的Docker容器的hosts文件已被修改，您可以使用新的主机名和IP地址在容器内部进行通信。