package main

import (
    db "../Database"
)

func main() {
    defer db.SqlDB.Close()
    
    router := initRouter()
    router.run(":8088")
}
