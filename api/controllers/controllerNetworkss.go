package controllers

import (
	"strconv"

	"github.com/msterzhang/onelist/api/database"
	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/api/repository"
	"github.com/msterzhang/onelist/api/repository/crud"

	"github.com/gin-gonic/gin"
)

func CreateNetworks(c *gin.Context) {
	networks := models.Networks{}
	err := c.ShouldBind(&networks)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错!", "data": networks})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositoryNetworkssCRUD(db)
	func(networksRepository repository.NetworksRepository) {
		networks, err := networksRepository.Save(networks)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "创建失败!", "data": networks})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "创建成功!", "data": networks})
	}(repo)
}

func DeleteNetworksById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryNetworkssCRUD(db)
	func(networksRepository repository.NetworksRepository) {
		networks, err := networksRepository.DeleteByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": networks})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "删除资源成功!", "data": networks})
	}(repo)
}

func UpdateNetworksById(c *gin.Context) {
	id := c.Query("id")
	networks := models.Networks{}
	err := c.ShouldBind(&networks)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错!", "data": networks})
		return
	}
	db := database.NewDb()
	repo := crud.NewRepositoryNetworkssCRUD(db)
	func(networksRepository repository.NetworksRepository) {
		networks, err := networksRepository.UpdateByID(id, networks)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": networks})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "更新资源成功!", "data": networks})
	}(repo)
}

func GetNetworksById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryNetworkssCRUD(db)
	func(networksRepository repository.NetworksRepository) {
		networks, err := networksRepository.FindByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": networks})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": networks})
	}(repo)
}

func GetNetworksList(c *gin.Context) {
	page, errPage := strconv.Atoi(c.Query("page"))
	size, errSize := strconv.Atoi(c.Query("size"))
	if errPage != nil {
		page = 1
	}
	if errSize != nil {
		size = 8
	}
	db := database.NewDb()
	repo := crud.NewRepositoryNetworkssCRUD(db)
	func(networksRepository repository.NetworksRepository) {
		networkss, num, err := networksRepository.FindAll(page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": networkss, "num": num})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": networkss, "num": num})
	}(repo)
}

func SearchNetworks(c *gin.Context) {
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
	repo := crud.NewRepositoryNetworkssCRUD(db)
	func(networksRepository repository.NetworksRepository) {
		networkss, num, err := networksRepository.Search(q, page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": networkss, "num": num})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": networkss, "num": num})
	}(repo)
}
