package routers

import (
  "github.com/gin-gonic/gin"
  . "GinLearn/GinLearn/apis"
 )
 
func InitRouter() *gin.Engine{
  router := gin.Default()
  //Hello World
  router.GET("/", IndexApi)
  //渲染html页面
  router.LoadHTMLGlob("views/*")
  router.GET("/home/index", ShowHtmlPage)
  //列表页面
  router.GET("/home/list", ListHtml)
  router.POST("/home/PageData", GetDataList)
  router.POST("/home/PageNextData", PageNextData)

  //新增页面
  router.GET("/home/add", AddHtml)
  router.POST("/home/saveadd", AddPersonApi)
  
   //编辑页面
   router.GET("/home/edit", EditHtml)
   router.POST("/home/saveedit", EditPersonApi)

    //删除
    router.POST("/home/delete", DeletePersonApi)
  return router
 }
 