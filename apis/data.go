package apis

import (
  "net/http"
  "github.com/gin-gonic/gin"
 . "GinLearn/GinLearn/models"
)
//Api调用的页面
func GetApiHtml(c *gin.Context){
	c.HTML(http.StatusOK,"api.html",gin.H{
		"title":"Go-Gin Api调用页面",
	})
}
//Json格式的数据
func GetJsonData(c *gin.Context) {  
  //得到请求的参数
  search:=c.PostForm("search")
  //得到用户的数据
  datalist:=GetPersonList(1,10,search)
  //得到记录的总数
  count:=GetRecordNum(search)
 //返回结果
 c.JSON(http.StatusOK, gin.H{
  "datalist": datalist,
  "count":count,
  "pagesize":3,
  "pageno":1,
 })
}

//Xml格式的数据
func GetXmlData(c *gin.Context) {  
	//得到请求的参数
	search:=c.PostForm("search")
	//得到用户的数据
	datalist:=GetPersonList(1,10,search)
	//得到记录的总数
	count:=GetRecordNum(search)
   //返回结果
   c.XML(http.StatusOK, gin.H{
	"datalist": datalist,
	"count":count,
	"pagesize":3,
	"pageno":1,
   })
}

//Xml格式的数据
func GetYamlData(c *gin.Context) {  
	//得到请求的参数
	search:=c.PostForm("search")
	//得到用户的数据
	datalist:=GetPersonList(1,10,search)
	//得到记录的总数
	count:=GetRecordNum(search)
   //返回结果
   c.YAML(http.StatusOK, gin.H{
	"datalist": datalist,
	"count":count,
	"pagesize":3,
	"pageno":1,
   })
  }

  //Json格式的数据
func GetParamsJsonData(c *gin.Context) {  
	//得到请求的参数
	search:=c.PostForm("search")
	//得到用户的数据
	datalist:=GetPersonList(1,10,search)
	//得到记录的总数
	count:=GetRecordNum(search)
   //返回结果
   c.JSON(http.StatusOK, gin.H{
	"datalist": datalist,
	"count":count,
	"pagesize":3,
	"pageno":1,
	"search":search,
   })
  }