package controllers

import (
	"fmt"
	"log"
	"html/template"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**内容页面**/
func Contenthtml(c *gin.Context){

    //模板文件的拼接
	t, err := template.ParseFiles("views/layout.html", "views/head.tpl", 
		"views/content.html","views/sidebar.tpl","views/scripts.tpl")
	//备注:参数1》模板页面；参数2》css部分；参数3》内容部分；
	//参数4》底部版权信息部分;参数5》页面中使用到的js部分
    if err != nil {
        log.Fatal(err)
    }
	fmt.Println(t)
	//渲染html文件
	c.HTML(http.StatusOK,"layout.html", gin.H{
		"title": "布局页面",
	})
}

