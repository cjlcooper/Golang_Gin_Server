package main

import (
  "fmt"
  "github.com/gin-gonic/gin"
  "net/http"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type User struct {
    UserName        string `json:"user_name" form:"user_name"`
    Password        string `json:"pass_word" form:"pass_word"`
}

func main() {
    fmt.Println("Start Server ...")
    InitDatabase()
    CreateRouter()
}

//Create Router
func CreateRouter() {
    // set mode
    //gin.SetMode(gin.ReleaseMode)
    gin.SetMode(gin.DebugMode)
    // Register Middleware
    router := gin.Default()
    router.Use(Middleware)
    // Test
    router.GET("/test/get_test",GetTest)
    router.POST("/test/post_test",PostTest)
    router.POST("/test/insert_test",InsertTest)
    // Listen Server
    //http.ListenAndServe(":8080",router)
    router.Run(":8088")
}

// Middle Ware
func Middleware(c *gin.Context) {
    fmt.Println("Start Middleware")
}

// Init Database
func InitDatabase() { 
    var err error
    db, err = sql.Open("mysql", "root:cjl1992@tcp(127.0.0.1:3306)/TEST?charset=utf8")
    //defer db.Close()
    if err != nil {
        fmt.Println(err)
    }

    db.SetMaxIdleConns(20)
    db.SetMaxOpenConns(20)

    if err := db.Ping(); err != nil {
        fmt.Println(err)
    }
}

// Router Function
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

//User func
func (user *User) AddUser() {
    rs, err := db.Exec("INSERT INTO test(user_name, pass_word) VALUES (?,?)",user.UserName,user.Password)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(rs)
    return
}

//get select result
//func (user *User)GetAllUser() (users []User, err error) {
//    users = make([]User, 0)
//    rows, err := db.Query("SELECT *FROM test")
//    if err != nil {
//        fmt.Println(err)
//    }
    
//    for rows.Next() {
//        fmt.Println(rows)
//    }
//    return nil
//}

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


