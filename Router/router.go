package main

import (
    "github.com/gin-gonic/gin"
    . "../Apis"
)

func initRouter() *gin.Engine {
    router := gin.Default()
    //router.Use(Middleware)
    // Test
    router.GET("/test/get_test",GetTest)
    router.POST("/test/post_test",PostTest)
    router.POST("/test/insert_test",InsertTest)
    return router
}
