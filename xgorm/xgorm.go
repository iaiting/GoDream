package main

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)


type User struct {
    gorm.Model
    Name    string
    Email   string
}


type Email struct {
    gorm.Model
    Email   string
    UserID  uint
}


func RdbGodreamEmail() {
    db, err := gorm.Open("postgres", "host=127.0.0.1 user=hengjun dbname=godream sslmode=disable password=123456")
    if err != nil {
        fmt.Println("数据库连接错误")
    }
    // db.SingularTable(true)
    db.LogMode(true)
    defer db.Close()

    db.CreateTable(&Email{})

    var email1 = Email {UserID:1, Email:"testUser1@163.com"}
    var email2 = Email {UserID:2, Email:"testUser2@163.com"}

    db.NewRecord(email1)
    db.Create(&email1)
    db.Create(&email2)
    db.NewRecord(email1)
}

func RdbGodreamUser() {
    db, err := gorm.Open("postgres", "host=127.0.0.1 user=hengjun dbname=godream sslmode=disable password=123456")
    if err != nil {
        fmt.Println("数据库连接错误")
    }
    db.SingularTable(true)
    db.LogMode(true)
    defer db.Close()

    db.CreateTable(&User{})

    var user1 = User {Name:"testUser1", Email:"testUser1@163.com"}
    var user2 = User {Name:"testUser1", Email:"testUser2@163.com"}

    db.NewRecord(user1)
    db.Create(&user1)
    db.Create(&user2)
    db.NewRecord(user1)
}

func rdb_init() {
    db, err := gorm.Open("postgres", "host=127.0.0.1 user=hengjun dbname=godream sslmode=disable password=123456")
    if err != nil {
        fmt.Println("数据库连接错误")
    }
    defer db.Close()
}

func main() {
    rdb_init()
    RdbGodreamEmail()
    RdbGodreamUser()
}
