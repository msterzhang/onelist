package initconfig

import (
	"fmt"
	"log"
	"os"

	"github.com/msterzhang/onelist/api"
	"github.com/msterzhang/onelist/api/database"
	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/api/security"
	"github.com/msterzhang/onelist/api/utils/tools"
	"github.com/msterzhang/onelist/config"
)

var configEnv = `
# 服务设置
# 注意要改为未被占用的端口
API_PORT=5245
FaviconicoUrl=https://wework.qpic.cn/wwpic/818353_fizV30xbQCGPQRP_1677394564/0
API_SECRET=%s

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
ImgUrl=https://image.tmdb.org
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
`

func InitConfigEnv() error {
	content := []byte(fmt.Sprintf(configEnv, tools.RandStringRunes(32)))
	err := os.WriteFile("config.env", content, 0644)
	if err != nil {
		return err
	}
	return nil
}

func AdminData() (models.User, error) {
	api.InitServer()
	db := database.NewDb()
	user := models.User{}
	err := db.Model(&models.User{}).Where("user_email = ?", config.UserEmail).First(&user).Error
	if err != nil {
		return models.User{}, err
	}
	pwd, err := security.DecodePassword(user.UserPassword)
	if err != nil {
		log.Fatal(err)
	}
	user.UserPassword = pwd
	return user, nil
}
