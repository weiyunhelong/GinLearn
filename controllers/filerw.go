package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
/**文件读写创建删除操作页面**/
func Filerwhtml(c *gin.Context){
	c.HTML(http.StatusOK, "filerw.html", gin.H{
		"title": "GIN: 文件读写创建删除操作布局页面",
	})
}

/**创建文件**/
func FilerCreate(c *gin.Context){
	c.HTML(http.StatusOK, "filerw.html", gin.H{
		"title": "GIN: 文件读写创建删除操作布局页面",
	})
}
/**将内容写入文件**/
func FilerWrite(c *gin.Context){
	c.HTML(http.StatusOK, "filerw.html", gin.H{
		"title": "GIN: 文件读写创建删除操作布局页面",
	})
}
/**读取文件内容**/
func FilerRead(c *gin.Context){
	c.HTML(http.StatusOK, "filerw.html", gin.H{
		"title": "GIN: 文件读写创建删除操作布局页面",
	})
}
/**删除文件**/
func FilerDelete(c *gin.Context){
	c.HTML(http.StatusOK, "filerw.html", gin.H{
		"title": "GIN: 文件读写创建删除操作布局页面",
	})
}