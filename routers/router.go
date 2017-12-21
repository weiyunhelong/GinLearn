package routers

import (
  "github.com/gin-gonic/gin"
  . "GinLearn/GinLearn/apis" //api部分
  . "GinLearn/GinLearn/controllers" //constroller部分
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

    //Bootstrap布局页面
    router.GET("/home/bootstrap", Bootstraphtml)

    //文件的上传和下载 
    router.GET("/home/fileopt", Fileopthtml)
    router.POST("/home/fileuplaod", Fileupload)
    router.GET("/home/filedown", Filedown)

    //文件的创建删除和读写
    router.GET("/home/filerw", Filerwhtml)
    router.POST("/home/addfile", FilerCreate)//创建文件
    router.POST("/home/writefile", FilerWrite)//写入文件
    router.POST("/home/readfile", FilerRead)//读取文件
    router.POST("/home/deletefile", FilerDelete)//删除文件
    
    return router
 }
 