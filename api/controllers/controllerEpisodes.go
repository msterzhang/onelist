package controllers

import (
	"strconv"

	"github.com/msterzhang/onelist/api/database"
	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/api/repository"
	"github.com/msterzhang/onelist/api/repository/crud"

	"github.com/gin-gonic/gin"
)

func CreateEpisode(c *gin.Context) {
	episode := models.Episode{}
	err := c.ShouldBind(&episode)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错!", "data": episode})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositoryEpisodesCRUD(db)
	func(episodeRepository repository.EpisodeRepository) {
		episode, err := episodeRepository.Save(episode)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "创建失败!", "data": episode})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "创建成功!", "data": episode})
	}(repo)
}

func DeleteEpisodeById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryEpisodesCRUD(db)
	func(episodeRepository repository.EpisodeRepository) {
		episode, err := episodeRepository.DeleteByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": episode})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "删除资源成功!", "data": episode})
	}(repo)
}

func UpdateEpisodeById(c *gin.Context) {
	id := c.Query("id")
	episode := models.Episode{}
	err := c.ShouldBind(&episode)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错!", "data": episode})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositoryEpisodesCRUD(db)
	func(episodeRepository repository.EpisodeRepository) {
		episode, err := episodeRepository.UpdateByID(id, episode)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": episode})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "更新资源成功!", "data": episode})
	}(repo)
}

func GetEpisodeById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryEpisodesCRUD(db)
	func(episodeRepository repository.EpisodeRepository) {
		episode, err := episodeRepository.FindByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": episode})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": episode})
	}(repo)
}

func GetEpisodeList(c *gin.Context) {
	page, errPage := strconv.Atoi(c.Query("page"))
	size, errSize := strconv.Atoi(c.Query("size"))
	if errPage != nil {
		page = 1
	}
	if errSize != nil {
		size = 8
	}
	db := database.NewDb()
	repo := crud.NewRepositoryEpisodesCRUD(db)
	func(episodeRepository repository.EpisodeRepository) {
		episodes, num, err := episodeRepository.FindAll(page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": episodes, "num": num})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": episodes, "num": num})
	}(repo)
}

func SearchEpisode(c *gin.Context) {
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
	repo := crud.NewRepositoryEpisodesCRUD(db)
	func(episodeRepository repository.EpisodeRepository) {
		episodes, num, err := episodeRepository.Search(q, page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": episodes, "num": num})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": episodes, "num": num})
	}(repo)
}
