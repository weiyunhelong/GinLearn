package controllers

import (
	"io/ioutil"
	"fmt"
	"os"
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
	iscreate:=true    //创建文件是否成功
	//创建文件
	f, err:= os.Create("static/txtfile/log.text")
	if err!=nil{
		iscreate=false		
	}
	defer f.Close()
	fmt.Println(f)
	//返回结果
	c.JSON(http.StatusOK, gin.H{
		"path":"static/txtfile/log.text",
		"success":iscreate,
	})
}
/**将内容写入文件**/
func FilerWrite(c *gin.Context){
	iswrite:=true                 //写入文件是否成功 
	//需要写入到文件的内容
	info:=c.PostForm("info")
	path:=c.PostForm("path")

	d1 := []byte(info)
	err := ioutil.WriteFile(path, d1, 0644)
	
    if err!=nil{
		iswrite=false    
	}
	//返回结果
	c.JSON(http.StatusOK, gin.H{
		"success":iswrite,
		"info":info,
	})
}
/**读取文件内容**/
func FilerRead(c *gin.Context){
	 isread:=true                 //读取文件是否成功 
	 path:=c.PostForm("path")
	 //文件读取任务是将文件内容读取到内存中。
	 info, err := ioutil.ReadFile(path)
	 if err!=nil{
		 fmt.Println(err)
		 isread=false  
	 }
	 fmt.Println(info)
	 result:=string(info)
	 
	//返回结果
	c.JSON(http.StatusOK, gin.H{
		"content":result,
		"success":isread,
	})
}
/**删除文件**/
func FilerDelete(c *gin.Context){
	
	isremove:=false                      //删除文件是否成功 
	path :=c.PostForm("path")    //源文件路径
	
	//删除文件
	cuowu := os.Remove(path)               

    if cuowu != nil {
        //如果删除失败则输出 file remove Error!
        fmt.Println("file remove Error!")
        //输出错误详细信息
        fmt.Printf("%s", cuowu)
    } else {
        //如果删除成功则输出 file remove OK!
		fmt.Print("file remove OK!")
		isremove=true
    }
	//返回结果
	c.JSON(http.StatusOK, gin.H{
		"success":isremove,
	})
}