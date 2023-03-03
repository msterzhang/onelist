package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/msterzhang/onelist/api/models"
)

var (
	EnvFile       = "config.env"
	PORT          = 0
	Title         = ""
	FaviconicoUrl = ""
	SECRETKEY     []byte
	DBDRIVER      = ""
	DBURL         = ""
	DBDATAURL     = ""
	DbName        = ""
	KeyDb         = ""
	UserEmail     = ""
	UserPassword  = ""
	DownLoadImage = ""
	ImgUrl        = ""
	VideoTypes    = ""
	UA            = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36"
)

// Load the server PORT
func Load() {
	var err error
	err = godotenv.Load(EnvFile)
	if err != nil {
		return
	}
	PORT, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		PORT = 9000
	}
	Env := os.Getenv("Env")
	Title = os.Getenv("Title")
	FaviconicoUrl = os.Getenv("FaviconicoUrl")
	if Env == "Debug" {
		DBDATAURL = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/?charset=utf8mb4", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD_Debug"))
		DBURL = fmt.Sprintf("%s:%s@/%s",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD_Debug"),
			os.Getenv("DB_NAME"),
		) + "?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"
	} else {
		DBDATAURL = fmt.Sprintf("%s:%s@/?charset=utf8mb4", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD_Release"))
		//当数据库为docker，注意替换：fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)
		DBURL = fmt.Sprintf("%s:%s@/%s",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD_Release"),
			os.Getenv("DB_NAME"),
		) + "?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"
	}
	DBDRIVER = os.Getenv("DB_DRIVER")
	SECRETKEY = []byte(os.Getenv("API_SECRET"))
	DbName = os.Getenv("DbName")
	KeyDb = os.Getenv("KeyDb")
	UserEmail = os.Getenv("UserEmail")
	UserPassword = os.Getenv("UserPassword")
	DownLoadImage = os.Getenv("DownLoadImage")
	ImgUrl = os.Getenv("ImgUrl")
	VideoTypes = os.Getenv("VideoTypes")
}

// 获取配置
func GetConfig() models.Config {
	config := models.Config{
		Title:         Title,
		DownLoadImage: DownLoadImage,
		ImgUrl:        ImgUrl,
		KeyDb:         KeyDb,
		FaviconicoUrl: FaviconicoUrl,
		VideoTypes:    VideoTypes,
	}
	return config
}

// 设置配置
func SetConfig(config models.Config) {
	Title = config.Title
	DownLoadImage = config.DownLoadImage
	ImgUrl = config.ImgUrl
	KeyDb = config.KeyDb
	FaviconicoUrl = config.FaviconicoUrl
	VideoTypes = config.VideoTypes
}

// 保存配置
func SaveConfig(config models.Config) (models.Config, error) {
	b, err := os.ReadFile(EnvFile)
	if err != nil {
		return models.Config{}, err
	}
	data := strings.ReplaceAll(string(b), "Title="+Title, "Title="+config.Title)
	data = strings.ReplaceAll(data, "DownLoadImage="+DownLoadImage, "DownLoadImage="+config.DownLoadImage)
	data = strings.ReplaceAll(data, "ImgUrl="+ImgUrl, "ImgUrl="+config.ImgUrl)
	data = strings.ReplaceAll(data, "FaviconicoUrl="+FaviconicoUrl, "FaviconicoUrl="+config.FaviconicoUrl)
	data = strings.ReplaceAll(data, "KeyDb="+KeyDb, "KeyDb="+config.KeyDb)
	data = strings.ReplaceAll(data, "VideoTypes="+VideoTypes, "VideoTypes="+config.VideoTypes)
	content := []byte(data)
	err = os.WriteFile(EnvFile, content, 0644)
	if err != nil {
		return models.Config{}, err
	}
	SetConfig(config)
	return GetConfig(), nil

}
