package Modles

import (
    "fmt"
    db "../Database"
)

type User struct {
    UserName        string `json:"user_name" form:"user_name"`
    Password        string `json:"pass_word" form:"pass_word"`
}

func (user *User) AddUser() {
    rs, err := db.SqlDB.Exec("INSERT INTO test(user_name, pass_word) VALUES (?,?)",user.UserName,user.Password)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(rs)
    return
}

func (user *User) GetAllUser() {
    rs, err := db.SqlDB.Exec("select * from test")
    defer rs.Close()

    if err != nil {
        fmt.Println(err)
    }


}

