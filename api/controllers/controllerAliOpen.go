package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/msterzhang/onelist/plugins/alist"
)

func AliOpenVideo(c *gin.Context) {
	aliOpenForm := alist.AliOpenForm{}
	err := c.ShouldBind(&aliOpenForm)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "表单解析出错!", "data": aliOpenForm})
		return
	}
	data, err := alist.AlistAliOpenVideo(aliOpenForm.File, aliOpenForm.GalleryUid)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": err.Error(), "data": ""})
	}
	c.JSON(200, gin.H{"code": 200, "msg": "success", "data": data.Data})
}
