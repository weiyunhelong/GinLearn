package routers

import (
  "github.com/gin-gonic/gin"
  . "GinLearn/GinLearn/apis"
 )
 
func InitRouter() *gin.Engine{
  router := gin.Default()
 
  router.GET("/", IndexApi)
 
  router.POST("/person", AddPersonApi)
 
  //router.GET("/persons", GetPersonsApi)
 
  //router.GET("/person/:id", GetPersonApi)
 
  //router.PUT("/person/:id", ModPersonApi)
 
  //router.DELETE("/person/:id", DelPersonApi)
  
  //渲染html页面
  router.LoadHTMLGlob("views/*")
  router.GET("/home/index", ShowHtmlPage)
  
  return router
 }
 