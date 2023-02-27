package controllers

import (
	"strconv"

	"github.com/msterzhang/onelist/api/database"
	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/api/repository"
	"github.com/msterzhang/onelist/api/repository/crud"

	"github.com/gin-gonic/gin"
)

func CreateGenre(c *gin.Context) {
	genre := models.Genre{}
	err := c.ShouldBind(&genre)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错!", "data": genre})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositoryGenresCRUD(db)
	func(genreRepository repository.GenreRepository) {
		genre, err := genreRepository.Save(genre)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "创建失败!", "data": genre})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "创建成功!", "data": genre})
	}(repo)
}

func DeleteGenreById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryGenresCRUD(db)
	func(genreRepository repository.GenreRepository) {
		genre, err := genreRepository.DeleteByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": genre})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "删除资源成功!", "data": genre})
	}(repo)
}

func UpdateGenreById(c *gin.Context) {
	id := c.Query("id")
	genre := models.Genre{}
	err := c.ShouldBind(&genre)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错!", "data": genre})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositoryGenresCRUD(db)
	func(genreRepository repository.GenreRepository) {
		genre, err := genreRepository.UpdateByID(id, genre)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": genre})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "更新资源成功!", "data": genre})
	}(repo)
}

func GetGenreById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryGenresCRUD(db)
	func(genreRepository repository.GenreRepository) {
		genre, err := genreRepository.FindByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": genre})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": genre})
	}(repo)
}

func GetGenreList(c *gin.Context) {
	page, errPage := strconv.Atoi(c.Query("page"))
	size, errSize := strconv.Atoi(c.Query("size"))
	if errPage != nil {
		page = 1
	}
	if errSize != nil {
		size = 8
	}
	db := database.NewDb()
	repo := crud.NewRepositoryGenresCRUD(db)
	func(genreRepository repository.GenreRepository) {
		genres, num, err := genreRepository.FindAll(page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": genres, "num": num})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": genres, "num": num})
	}(repo)
}

func SearchGenre(c *gin.Context) {
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
	repo := crud.NewRepositoryGenresCRUD(db)
	func(genreRepository repository.GenreRepository) {
		genres, num, err := genreRepository.Search(q, page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": genres, "num": num})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": genres, "num": num})
	}(repo)
}
