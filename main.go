package main

import (
  "fmt"
  "os"
  "github.com/joho/godotenv"
  "github.com/jinzhu/gorm"
  _ "github.com/go-sql-driver/mysql"
  "time"
)

type User struct {
    Id int64 `gorm:"primary_key"`
    Name string `sql:"size:255"`
    CreatedAt time.Time
    UpdatedAT time.Time
    DeletedAt time.Time
}

func GetDBConn() *gorm.DB {
   db, err := gorm.Open(GetDBConfig())
   if err != nil {
      panic(err)
   }

   db.LogMode(true)
   return db
}

func GetDBConfig() (string, string) {
   DBMS := "mysql"
   USER := os.Getenv("USER")
   PASS := os.Getenv("PASS")
   PROTOCOL := ""
   DBNAME := os.Getenv("DBNAME")
   OPTION := "charset=utf8&parseTime=True&loc=Local"

   CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?" + OPTION

   return DBMS, CONNECT
}

func main() {
  err := godotenv.Load(fmt.Sprintf("../%s.env", os.Getenv("GO_ENV")))
    if err != nil {
        // .env読めなかった場合の処理
    }

    env := os.Getenv("GO_ENV")
    DBUser := os.Getenv("USER")
    DBPass := os.Getenv("PASS")
    DBName := os.Getenv("DBNAME")

    fmt.Println(env)
    fmt.Println(DBUser)
    fmt.Println(DBName)
    fmt.Println(DBPass)

   db := GetDBConn()
   db.AutoMigrate(&User{})

}
