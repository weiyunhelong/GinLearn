package main

import (
 db "GinLearn/GinLearn/database"
 router "GinLearn/GinLearn/routers"
)

func main() {
   //数据库
   defer db.SqlDB.Close()
   //路由部分
   router:=router.InitRouter()
   //静态资源
   router.Static("/static", "./static")
   //运行的端口
   router.Run(":8000")
}
