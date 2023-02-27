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

func CreateStar(c *gin.Context) {
	star := models.Star{}
	err := c.ShouldBind(&star)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错!", "data": star})
		return
	}
	if len(star.UserId) == 0 {
		star.UserId = c.GetString("UserId")
	}
	db := database.NewDb()
	repo := crud.NewRepositoryStarsCRUD(db)
	func(starRepository repository.StarRepository) {
		star, err := starRepository.Save(star)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "创建失败!", "data": star})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "创建成功!", "data": star})
	}(repo)
}

func DeleteStarById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryStarsCRUD(db)
	func(starRepository repository.StarRepository) {
		star, err := starRepository.DeleteByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": star})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "删除资源成功!", "data": star})
	}(repo)
}

func UpdateStarById(c *gin.Context) {
	id := c.Query("id")
	star := models.Star{}
	err := c.ShouldBind(&star)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错!", "data": star})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositoryStarsCRUD(db)
	func(starRepository repository.StarRepository) {
		star, err := starRepository.UpdateByID(id, star)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": star})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "更新资源成功!", "data": star})
	}(repo)
}

func ReNewStarByStar(c *gin.Context) {
	star := models.Star{}
	err := c.ShouldBind(&star)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错!", "data": star})
		return
	}
	if len(star.UserId) == 0 {
		star.UserId = c.GetString("UserId")
	}
	db := database.NewDb()
	repo := crud.NewRepositoryStarsCRUD(db)
	func(starRepository repository.StarRepository) {
		star, err := starRepository.ReNewStarByStar(star)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "处理失败!", "data": star})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "处理成功!", "data": star})
	}(repo)
}
func GetStarById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryStarsCRUD(db)
	func(starRepository repository.StarRepository) {
		star, err := starRepository.FindByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": star})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": star})
	}(repo)
}

func GetStarList(c *gin.Context) {
	page, errPage := strconv.Atoi(c.Query("page"))
	size, errSize := strconv.Atoi(c.Query("size"))
	if errPage != nil {
		page = 1
	}
	if errSize != nil {
		size = 8
	}
	db := database.NewDb()
	repo := crud.NewRepositoryStarsCRUD(db)
	func(starRepository repository.StarRepository) {
		stars, num, err := starRepository.FindAll(page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": stars, "num": num})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": stars, "num": num})
	}(repo)
}

func SearchStar(c *gin.Context) {
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
	repo := crud.NewRepositoryStarsCRUD(db)
	func(starRepository repository.StarRepository) {
		stars, num, err := starRepository.Search(q, page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": stars, "num": num})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": stars, "num": num})
	}(repo)
}

func GetStarDataList(c *gin.Context) {
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
	star := models.Star{UserId: UserId, DataType: dataType}
	db := database.NewDb()
	repo := crud.NewRepositoryStarsCRUD(db)
	func(starRepository repository.StarRepository) {
		stars, num, err := starRepository.FindAllByUser(star, page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": stars, "num": num})
			return
		}
		if dataType == "tv" {
			thetvs := []models.TheTv{}
			for _, starDb := range stars {
				thetv := models.TheTv{}
				err = db.Model(&models.TheTv{}).Where("id = ?", starDb.DataId).First(&thetv).Error
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
		for _, starDb := range stars {
			themovie := models.TheMovie{}
			err = db.Model(&models.TheMovie{}).Where("id = ?", starDb.DataId).First(&themovie).Error
			if err != nil {
				continue
			}
			themovie = service.TheMovieService(themovie, UserId)
			themovies = append(themovies, themovie)
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": themovies, "num": num})
	}(repo)
}
