package apis

import (
 "net/http"
 "log"
 "fmt"
 "github.com/gin-gonic/gin"
 . "GinLearn/GinLearn/models"
)
//初始页面
func IndexApi(c *gin.Context) {
  c.String(http.StatusOK, "Hello World!")
}

//新增记录
func AddPersonApi(c *gin.Context) {
   firstName := c.Request.FormValue("first_name")
   lastName := c.Request.FormValue("last_name")

   p := Person{FirstName: firstName, LastName: lastName}

   ra, err := p.AddPerson()
    if err != nil {
    log.Fatalln(err)
  }
   msg := fmt.Sprintf("insert successful %d", ra)
   c.JSON(http.StatusOK, gin.H{
     "msg": msg,
 })
}

//渲染html页面
func ShowHtmlPage(c *gin.Context) {
    c.HTML(http.StatusOK, "index.html", gin.H{})
  }
