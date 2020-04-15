package main

import (
  "strconv"
	"github.com/gin-gonic/gin"
  "github.com/jinzhu/gorm"
  _ "github.com/go-sql-driver/mysql"
)

type Himajin struct {
    gorm.Model
    Name   string
    Begintime string
    FinishTime string
    Todo string
}

//DBマイグレート
func dbInit() {
    db, err := gorm.Open("mysql", "root:@(localhost)/gopractice?charset=utf8&parseTime=True")
    if err != nil {
        panic("データベース失敗（dbInit）")
    }
    db.AutoMigrate(&Himajin{})
    defer db.Close()
}
//DB追加
func dbInsert(name string, begintime string, finishtime string, todo string) {
    db, err := gorm.Open("mysql", "root:@(localhost)/gopractice?charset=utf8&parseTime=True")
    if err != nil {
        panic("データベース失敗（dbInsert)")
    }
    db.Create(&Himajin{Name: name, Begintime: begintime, FinishTime: finishtime, Todo: todo})
    defer db.Close()
}

//DB更新
func dbUpdate(id int, name string, begintime string, finishtime string, todo string) {
    db, err := gorm.Open("mysql", "root:@(localhost)/gopractice?charset=utf8&parseTime=True")
    if err != nil {
        panic("データベース失敗（dbUpdate)")
    }
    var himajin Himajin
    db.First(&himajin, id)
    himajin.Name = name
    himajin.Begintime = begintime
    himajin.FinishTime = finishtime
    himajin.Todo = todo
    db.Save(&himajin)
    db.Close()
}

//DB削除
func dbDelete(id int) {
    db, err := gorm.Open("mysql", "root:@(localhost)/gopractice?charset=utf8&parseTime=True")
    if err != nil {
        panic("データベース失敗（dbDelete)")
    }
    var himajin Himajin
    db.First(&himajin, id)
    db.Delete(&himajin)
    db.Close()
}

//DB全取得
func dbGetAll() []Himajin {
    db, err := gorm.Open("mysql", "root:@(localhost)/gopractice?charset=utf8&parseTime=True")
    if err != nil {
        panic("データベース失敗(dbGetAll())")
    }
    var himajins []Himajin
    db.Order("created_at desc").Find(&himajins)
    db.Close()
    return himajins
}

//DB一つ取得
func dbGetOne(id int) Himajin {
    db, err := gorm.Open("mysql", "root:@(localhost)/gopractice?charset=utf8&parseTime=True")
    if err != nil {
        panic("データベース失敗(dbGetOne())")
    }
    var himajin Himajin
    db.First(&himajin, id)
    db.Close()
    return himajin
}





func main() {
    router := gin.Default()
    router.LoadHTMLGlob("./templates/*.html")

    dbInit()

    //Index
    router.GET("/", func(ctx *gin.Context) {
        himajin := dbGetAll()
        ctx.HTML(200, "index.html", gin.H{
            "himajins": himajin,
        })
    })

    //Create
    router.POST("/new", func(ctx *gin.Context) {
        name := ctx.PostForm("name")
        begintime := ctx.PostForm("begintime")
        finishtime := ctx.PostForm("finishtime")
        todo := ctx.PostForm("todo")
        dbInsert(name, begintime, finishtime, todo)
        ctx.Redirect(302, "/")
    })

    //Detail
    router.GET("/detail/:id", func(ctx *gin.Context) {
        n := ctx.Param("id")
        id, err := strconv.Atoi(n)
        if err != nil {
            panic(err)
        }
        himajin := dbGetOne(id)
        ctx.HTML(200, "detail.html", gin.H{"himajins": himajin})
    })

    //Update
    router.POST("/update/:id", func(ctx *gin.Context) {
        n := ctx.Param("id")
        id, err := strconv.Atoi(n)
        if err != nil {
            panic("ERROR")
        }
        name := ctx.PostForm("name")
        begintime := ctx.PostForm("begintime")
        finishtime := ctx.PostForm("finishtime")
        todo := ctx.PostForm("todo")
        dbUpdate(id, name, begintime, finishtime, todo)
        ctx.Redirect(302, "/")
    })

    //削除確認
    router.GET("/delete_check/:id", func(ctx *gin.Context) {
        n := ctx.Param("id")
        id, err := strconv.Atoi(n)
        if err != nil {
            panic("ERROR")
        }
        himajin := dbGetOne(id)
        ctx.HTML(200, "delete.html", gin.H{"himajins": himajin})
    })

    //Delete
    router.POST("/delete/:id", func(ctx *gin.Context) {
        n := ctx.Param("id")
        id, err := strconv.Atoi(n)
        if err != nil {
            panic("ERROR")
        }
        dbDelete(id)
        ctx.Redirect(302, "/")

    })

    router.Run()
}
