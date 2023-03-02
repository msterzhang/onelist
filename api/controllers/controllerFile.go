package controllers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func ImgServer(c *gin.Context) {
	path := c.Param("path")
	filePath := "images/" + path
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
