package progress

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/msterzhang/onelist/api/database"
	"github.com/msterzhang/onelist/api/models"
	"gorm.io/gorm"
)

func Get(c *gin.Context) {
	db := database.NewDb()
	UserId := c.Request.Header.Get("UserId")
	if UserId == "" {
		c.JSON(400, gin.H{"msg": "未获取到所需字段"})
		return
	}
	err := db.Model(&models.User{}).Where("UserId = ?", UserId).First(models.User{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(400, gin.H{"msg": "该用户不存在"})
		return
	}
	progress := &models.Progress{}
	err = db.Model(&models.Progress{}).Where("user_id = ?", UserId).Find(progress).Error
	if err != nil {
		c.JSON(200, gin.H{})
		return
	}
	c.JSON(200, gin.H{"data": progress.Data})
	return
}

func Update(c *gin.Context) {
	db := database.NewDb()
	UserId := c.Request.Header.Get("UserId")
	if UserId == "" {
		c.JSON(400, gin.H{"msg": "未获取到所需字段"})
		return
	}
	err := db.Model(&models.User{}).Where("user_id = ?", UserId).First(&models.User{}).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(400, gin.H{"msg": "该用户不存在"})
		return
	}
	requestBody := &models.ProgressRequestBody{}

	err = c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.JSON(400, gin.H{"msg": "缺少参数"})
		return
	}
	progress := &models.Progress{}
	dbprogress := &models.Progress{}
	err = db.Model(&models.Progress{}).Where("user_id = ?", UserId).First(progress).Error
	d, _ := json.Marshal(requestBody.Data)
	progress.Data = string(d)
	progress.UserId = UserId
	if errors.Is(err, gorm.ErrRecordNotFound) {
		progress.CreatedAt = dbprogress.CreatedAt
		err = db.Debug().Model(&models.Progress{}).Create(&progress).Error
		if err != nil {
			c.JSON(500, gin.H{"msg": "记录创建失败"})
			return
		}
		c.JSON(200, gin.H{"msg": "记录创建成功"})
		return
	}

	err = db.Model(&models.Progress{}).Where("id = ?", progress.Id).Select("*").Updates(&progress).Error
	if err != nil {
		c.JSON(500, gin.H{"msg": "记录更新失败"})
		return
	}
	c.JSON(200, gin.H{"msg": "记录更新成功"})
	return
}
