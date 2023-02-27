package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	PORT          = 0
	FaviconicoUrl = ""
	SECRETKEY     []byte
	DBDRIVER      = ""
	DBURL         = ""
	DBDATAURL     = ""
	DbName        = ""
	KeyDb         = ""
	UA            = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36"
)

// Load the server PORT
func init() {
	var err error
	err = godotenv.Load("config.env")
	if err != nil {
		log.Fatal(err)
	}
	PORT, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		PORT = 9000
	}
	Env := os.Getenv("Env")
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
	
}
