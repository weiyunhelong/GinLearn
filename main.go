package main

import (
 db "GinLearn/GinLearn/database"
 router "GinLearn/GinLearn/routers"
)

func main() {
 defer db.SqlDB.Close()
 router:=router.InitRouter()
 router.Run(":8000")
}
