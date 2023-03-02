package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/config"
)

func GetConfig(c *gin.Context) {
	configData := config.GetConfig()
	c.JSON(200, gin.H{"code": 200, "msg": "获取成功!", "data": configData})
}

func SaveConfig(c *gin.Context) {
	configData := models.Config{}
	err := c.ShouldBind(&configData)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": err.Error(), "data": ""})
		return
	}
	data, err := config.SaveConfig(configData)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": err.Error(), "data": ""})
		return
	}
	c.JSON(200, gin.H{"code": 200, "msg": "保存成功!", "data": data})
}
