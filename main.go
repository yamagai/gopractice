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
    models.UserDbInit()
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

    //以下seesion練習
    user := router.Group("/user")
    {
        user.POST("/signup", routes.UserSignUp)
        user.POST("/login", routes.UserLogIn)
    }
    router.GET("/login", routes.LogIn)
    router.GET("/signup", routes.SignUp)
    router.NoRoute(routes.NoRoute)

    router.Run()
}
