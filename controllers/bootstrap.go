package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
/**Bootstrap布局页面**/
func Bootstraphtml(c *gin.Context){
	c.HTML(http.StatusOK, "bootstrap.html", gin.H{
		"title": "GIN: Bootstrap布局页面",
	})
}