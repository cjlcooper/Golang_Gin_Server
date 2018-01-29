package Apis

import (
    "net/http"
    "fmt"
    . "../Modles"
    "github.com/gin-gonic/gin"
)

func IndexApi(c *gin.Context) {
    c.String(http.StatusOK, "It works")
}

func GetTest(c *gin.Context) {
    value, exist := c.GetQuery("key")
    if !exist {
        value = "the key is not exist!"
    }
    c.Data(http.StatusOK, "text/plain", []byte(fmt.Sprintf("get success! %s\n", value)))
    return
}

func PostTest(c *gin.Context) {
    type JsonHolder struct {
        Id   int    `json:"id"`
        Name string `json:"name"`
    }
    holder := JsonHolder{Id: 1, Name: "my name"}
    c.JSON(http.StatusOK, holder)
    return
}

//Insert User Data
func InsertTest(c *gin.Context) {
    userName := c.Request.FormValue("userName")
    password := c.Request.FormValue("password")
   
    user := User{UserName:userName, Password:password}
     
    user.AddUser()
    c.JSON(http.StatusOK, gin.H{
        "msg":"success",
    })
    return
}

//Get User Data
func GetUserData(c *gin.Context) {

}
