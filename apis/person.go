package apis

import (
	"fmt"
	"strconv"
  "net/http"
  "log"
  "github.com/gin-gonic/gin"
 . "GinLearn/GinLearn/models"
)
//初始页面
func IndexApi(c *gin.Context) {
  c.String(http.StatusOK, "Hello World!")
}
//渲染html页面
func ShowHtmlPage(c *gin.Context) {
  c.HTML(http.StatusOK, "index.html", gin.H{
      "title": "GIN: HTML页面",
  })
}


//列表页面
func ListHtml(c *gin.Context) {  
  c.HTML(http.StatusOK, "list.html", gin.H{
      "title": "GIN: 用户列表页面",
  })
}
//列表页面数据
func GetDataList(c *gin.Context) {  
  //得到请求的参数
  search:=c.PostForm("search")
  num:=c.PostForm("pageno")
  pageno,err:= strconv.Atoi(num) 
  if err!=nil{
    log.Fatalln(err)
  }
  //得到数据集
  datalist:=GetPersonList(pageno,3,search)

  //得到记录的总数
  count:=GetRecordNum(search)
 //返回结果
 c.JSON(http.StatusOK, gin.H{
  "datalist": datalist,
  "count":count,
  "pagesize":3,
  "pageno":pageno,
 })
}
//列表页面数据
func PageNextData(c *gin.Context) {  
  //得到请求的参数
  search:=c.PostForm("search")
  num:=c.PostForm("pageno")
  pageno,err:= strconv.Atoi(num) 
  if err!=nil{
    log.Fatalln(err)
  }
  //得到数据集
  datalist:=GetPersonList(pageno,3,search)

  //得到记录的总数
  count:=GetRecordNum(search)
 //返回结果
 c.JSON(http.StatusOK, gin.H{
  "datalist": datalist,
  "count":count,
  "pagesize":3,
  "pageno":pageno,
 })
}
//新增页面
func AddHtml(c *gin.Context){
  c.HTML(http.StatusOK, "add.html", gin.H{
    "title": "GIN: 新增用户页面",
  })
}
//新增记录
func AddPersonApi(c *gin.Context) {
 
   //得到请求的参数
   firstName := c.PostForm("first_name")
   lastName := c.PostForm("last_name")
 
   //赋值
   p := Person{FirstName: firstName, LastName: lastName}
   //调用模型中的新增的方法
   ra:= p.AddPerson()
   //返回结果
   c.JSON(http.StatusOK, gin.H{
    "success": ra,
  })
}
//编辑页面
func EditHtml(c *gin.Context){
  //得到URL请求的参数
  num:=c.Query("id")

  id,err:= strconv.Atoi(num) 
 
  if err!=nil{
    log.Fatalln(err)
  }

  p:=GetPersonById(id)
   if p==nil{
    fmt.Println("得到数据错误")
   }else{
    fmt.Println(p)
    fmt.Println("得到数据正确")
   }

  c.HTML(http.StatusOK, "edit.html", gin.H{
    "title": "GIN: 编辑用户页面",
    "id":p.Id,
    "firstname":p.FirstName,
    "lastname":p.LastName,
  })
}
//编辑记录
func EditPersonApi(c *gin.Context) {
  fmt.Println("执行到此处")
  //得到请求的参数
  num:=c.PostForm("id")
  fmt.Println(num)
  id,err:= strconv.Atoi(num) 
  if err!=nil{
    log.Fatalln(err)
  }
  firstName := c.PostForm("first_name")
  lastName := c.PostForm("last_name")
  //赋值
  p := GetPersonById(id)
  p.FirstName=firstName
  p.LastName=lastName
  //调用模型中的编辑的方法
  ra:= p.EditPerson()
  
  //返回结果
  c.JSON(http.StatusOK, gin.H{
    "success": ra,
  })
}

//删除记录
func DeletePersonApi(c *gin.Context) {
 
  //得到请求的参数
   num:=c.PostForm("id")
   fmt.Println(num)
   id,err:= strconv.Atoi(num) 
   if err!=nil{
    log.Fatalln(err)
  }
  //调用模型中的删除的方法
  ra:= DeletePerson(id)
  if ra == false {
   log.Fatalln("删除失败")
  }
  //返回结果
  c.JSON(http.StatusOK, gin.H{
    "success": ra,
  })
}

