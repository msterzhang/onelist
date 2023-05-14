package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/msterzhang/onelist/api/utils/dir"
	"github.com/msterzhang/onelist/api/utils/tools"
)

// 图片文件服务
func ImgServer(c *gin.Context) {
	path := c.Param("path")
	filePath := "images" + path
	c.Writer.WriteHeader(200)
	b, err := os.ReadFile(filePath)
	if err != nil {
		c.Writer.WriteHeader(http.StatusNotFound)
		c.Writer.Flush()
		return
	}
	_, err = c.Writer.Write(b)
	if err != nil {
		c.Writer.WriteHeader(http.StatusNoContent)
		c.Writer.Flush()
		return
	}
	c.Writer.Header().Add("Content-Type", "image/*")
	c.Writer.Flush()
}

// 本地文件服务
func FileServer(c *gin.Context) {
	file := c.Param("path")
	if len(file) < 1 {
		c.String(http.StatusBadRequest, "文件不存在!")
		return
	}
	file = file[1:]
	if !dir.FileExists(file) {
		c.String(http.StatusBadRequest, "文件不存在!")
		return
	}
	fileName := filepath.Base(file)
	c.Header("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, fileName))
	c.Header("Content-Type", "application/octet-stream")
	c.File(file)
}

func GalleryImgServer(c *gin.Context) {
	path := c.Param("path")
	filePath := "./images" + path
	c.Writer.WriteHeader(200)
	b, err := os.ReadFile(filePath)
	if err != nil {
		c.Writer.WriteHeader(http.StatusNotFound)
		c.Writer.Flush()
		return
	}
	_, err = c.Writer.Write(b)
	if err != nil {
		c.Writer.WriteHeader(http.StatusNoContent)
		c.Writer.Flush()
		return
	}
	c.Writer.Header().Add("Content-Type", "image/*")
	c.Writer.Flush()
}

func FileUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "没有获得文件!")
		return
	}
	id := tools.RandStringRunes(16)
	dst := "./images/w355_and_h200_multi_faces/" + id + path.Ext(file.Filename)
	data := "/gallery/w355_and_h200_multi_faces/" + id + path.Ext(file.Filename)
	c.SaveUploadedFile(file, dst)
	c.JSON(200, gin.H{"code": 200, "msg": "上传成功!", "data": data})
}
