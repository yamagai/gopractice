package main

import (
  "gopractice/models"
  "gopractice/routes"
	"github.com/gin-gonic/gin"
  _ "github.com/go-sql-driver/mysql"
)

func main() {
    router := gin.Default()
    router.LoadHTMLGlob("./templates/*.html")
    router.Static("/assets", "./assets")
    models.DbInit()
    //Index
    router.GET("/", routes.Index)
    //Create
    router.POST("/new", routes.Create)
    //Detail
    router.GET("/detail/:id", routes.Detail)
    //Update
    router.POST("/update/:id", routes.Update)
    //削除確認
    router.GET("/delete_check/:id", routes.Deletecheck)
    //Delete
    router.POST("/delete/:id", routes.Delete)
    router.Run()
}
