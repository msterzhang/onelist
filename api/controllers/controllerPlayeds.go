package controllers

import (
	"strconv"

	"github.com/msterzhang/onelist/api/database"
	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/api/repository"
	"github.com/msterzhang/onelist/api/repository/crud"
	"github.com/msterzhang/onelist/api/service"

	"github.com/gin-gonic/gin"
)

func CreatePlayed(c *gin.Context) {
	played := models.Played{}
	err := c.ShouldBind(&played)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错!", "data": played})
		return
	}
	if len(played.UserId) == 0 {
		played.UserId = c.GetString("UserId")
	}
	db := database.NewDb()
	repo := crud.NewRepositoryPlayedsCRUD(db)
	func(playedRepository repository.PlayedRepository) {
		played, err := playedRepository.Save(played)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "创建失败!", "data": played})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "创建成功!", "data": played})
	}(repo)
}

func DeletePlayedById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryPlayedsCRUD(db)
	func(playedRepository repository.PlayedRepository) {
		played, err := playedRepository.DeleteByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": played})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "删除资源成功!", "data": played})
	}(repo)
}

func ReNewPlayedByPlayed(c *gin.Context) {
	played := models.Played{}
	err := c.ShouldBind(&played)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错!", "data": played})
		return
	}
	if len(played.UserId) == 0 {
		played.UserId = c.GetString("UserId")
	}
	db := database.NewDb()
	repo := crud.NewRepositoryPlayedsCRUD(db)
	func(playedRepository repository.PlayedRepository) {
		played, err := playedRepository.ReNewByPlayed(played)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": played})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "处理成功!", "data": played})
	}(repo)
}

func UpdatePlayedById(c *gin.Context) {
	id := c.Query("id")
	played := models.Played{}
	err := c.ShouldBind(&played)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错!", "data": played})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositoryPlayedsCRUD(db)
	func(playedRepository repository.PlayedRepository) {
		played, err := playedRepository.UpdateByID(id, played)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": played})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "更新资源成功!", "data": played})
	}(repo)
}

func GetPlayedById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryPlayedsCRUD(db)
	func(playedRepository repository.PlayedRepository) {
		played, err := playedRepository.FindByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": played})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": played})
	}(repo)
}

func GetPlayedList(c *gin.Context) {
	page, errPage := strconv.Atoi(c.Query("page"))
	size, errSize := strconv.Atoi(c.Query("size"))
	if errPage != nil {
		page = 1
	}
	if errSize != nil {
		size = 8
	}
	db := database.NewDb()
	repo := crud.NewRepositoryPlayedsCRUD(db)
	func(playedRepository repository.PlayedRepository) {
		playeds, num, err := playedRepository.FindAll(page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": playeds, "num": num})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": playeds, "num": num})
	}(repo)
}

func SearchPlayed(c *gin.Context) {
	q := c.Query("q")
	if len(q) == 0 {
		c.JSON(200, gin.H{"code": 201, "msg": "参数错误!", "data": ""})
		return
	}
	page, errPage := strconv.Atoi(c.Query("page"))
	size, errSize := strconv.Atoi(c.Query("size"))
	if errPage != nil {
		page = 1
	}
	if errSize != nil {
		size = 8
	}
	db := database.NewDb()
	repo := crud.NewRepositoryPlayedsCRUD(db)
	func(playedRepository repository.PlayedRepository) {
		playeds, num, err := playedRepository.Search(q, page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": playeds, "num": num})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": playeds, "num": num})
	}(repo)
}

func GetPlayedDataList(c *gin.Context) {
	dataType := c.Query("data_type")
	if len(dataType) == 0 {
		c.JSON(200, gin.H{"code": 201, "msg": "参数错误!", "data": ""})
		return
	}
	page, errPage := strconv.Atoi(c.Query("page"))
	size, errSize := strconv.Atoi(c.Query("size"))
	if errPage != nil {
		page = 1
	}
	if errSize != nil {
		size = 8
	}
	UserId := c.GetString("UserId")
	played := models.Played{UserId: UserId, DataType: dataType}
	db := database.NewDb()
	repo := crud.NewRepositoryPlayedsCRUD(db)
	func(playedRepository repository.PlayedRepository) {
		playeds, num, err := playedRepository.FindAllByUser(played, page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": playeds, "num": num})
			return
		}
		if dataType == "tv" {
			thetvs := []models.TheTv{}
			for _, playedDb := range playeds {
				thetv := models.TheTv{}
				err = db.Model(&models.TheTv{}).Where("id = ?", playedDb.DataId).First(&thetv).Error
				if err != nil {
					continue
				}
				thetv = service.TheTvService(thetv, UserId)
				thetvs = append(thetvs, thetv)
			}
			c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": thetvs, "num": num})
			return
		}
		themovies := []models.TheMovie{}
		for _, playedDb := range playeds {
			themovie := models.TheMovie{}
			err = db.Model(&models.TheMovie{}).Where("id = ?", playedDb.DataId).First(&themovie).Error
			if err != nil {
				continue
			}
			themovie = service.TheMovieService(themovie, UserId)
			themovies = append(themovies, themovie)
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": themovies, "num": num})
	}(repo)
}
