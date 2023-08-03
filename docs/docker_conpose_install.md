### 1.主机新建一个文件夹：

/root/onelist

文件夹中创建一个`/root/onelist/docker-compose.yml`文件，写入以下内容
```
version: '3.3'
services:
  onelist:
    restart: always
    container_name: onelist
    image: 'msterzhang/onelist:latest'
    volumes:
      - '/root/onelist/config:/config'
    ports:
      - '5245:5245'
    environment:
      - PUID=0
      - PGID=0
      - UMASK=022
      - TZ=Asia/Shanghai
    extra_hosts:
      - 'api.themoviedb.org:13.224.161.90'
      - 'api.themoviedb.org:13.35.67.86'
      - 'api.themoviedb.org:13.249.175.212'
      - 'api.themoviedb.org:13.35.161.120'
      - 'image.themoviedb.org:104.16.61.155'
      - 'www.themoviedb.org:54.192.151.79'

```


### 2.运行命令启动项目容器
```
# 切换到docker-compose.yml所在目录
cd /root/onelist

# 启动项目
docker-compose up -d
```

### 3.修改配置
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
### 4.重启项目容器
```
docker-compose up -d
```

### 5.监控容器，有最新版onelist，自动下载安装
```
docker pull containrrr/watchtower

docker run --rm -v /var/run/docker.sock:/var/run/docker.sock containrrr/watchtower -c --run-once onelist
```
> 注意：进入后台后需要删除初始化的xxxx.@qq.com账号，防止被人登录