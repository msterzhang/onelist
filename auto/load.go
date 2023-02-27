package auto

import (
	"database/sql"
	"errors"
	"log"

	"github.com/msterzhang/onelist/api/database"
	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/config"
	"gorm.io/gorm"
)

func init() {
	if config.DBDRIVER == "mysql" {
		db, err := sql.Open(config.DBDRIVER, config.DBDATAURL)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		_, err = db.Exec("CREATE DATABASE " + config.DbName + " default character set utf8mb4 collate utf8mb4_general_ci")
		if err != nil {
			log.Println("数据库已存在!")
			InitDatabase()
			return
		}
		log.Println("数据库创建成功!", err)
	}
	InitDatabase()
}

func InitAmdin() {
	db := database.NewDb()
	user := models.User{}
	err := db.Model(&models.User{}).Where("user_email = ?", config.UserEmail).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		user.UserEmail = config.UserEmail
		user.UserPassword = config.UserPassword
		user.IsAdmin = true
		db.Model(&models.User{}).Create(&user)
	}
	if err != nil {
		log.Fatal(err)
	}
}

func InitDatabase() {
	err := database.InitDb()
	if err != nil {
		log.Fatal("Gorm初始化数据库失败!报错：" + err.Error())
	}
	InitAmdin()
}

func Load() {
	var err error
	db := database.NewDb()
	err = db.AutoMigrate(&models.TheTv{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&models.Season{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&models.LastEpisodeToAir{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&models.NextEpisodeToAir{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&models.Networks{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&models.Genre{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&models.ProductionCompanie{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&models.ProductionCountrie{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&models.SpokenLanguage{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&models.BelongsToCollection{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&models.TheMovie{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&models.TheCredit{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&models.CastItem{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&models.CrewItem{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&models.Episode{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&models.TheSeason{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&models.ThePerson{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&models.Work{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&models.Gallery{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&models.ErrFile{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&models.Star{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&models.Heart{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&models.Played{})
	if err != nil {
		log.Fatal(err)
	}
}
