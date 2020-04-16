package routes

import (
  "gopractice/models"
  "net/http"
  "strconv"
  "github.com/gin-gonic/gin"
)

//index
func Index(ctx *gin.Context) {
    himajin := models.DbGetAll()
    ctx.HTML(http.StatusOK, "index.html", gin.H{
        "himajins": himajin,
    })
}

//create
func Create(ctx *gin.Context) {
    name := ctx.PostForm("name")
    begintime := ctx.PostForm("begintime")
    finishtime := ctx.PostForm("finishtime")
    todo := ctx.PostForm("todo")
    models.DbInsert(name, begintime, finishtime, todo)
    ctx.Redirect(302, "/")
}

//Detail
func Detail(ctx *gin.Context) {
    n := ctx.Param("id")
    id, err := strconv.Atoi(n)
    if err != nil {
        panic(err)
    }
    himajin := models.DbGetOne(id)
    ctx.HTML(http.StatusOK, "detail.html", gin.H{"himajins": himajin})
}

//update
func Update(ctx *gin.Context) {
    n := ctx.Param("id")
    id, err := strconv.Atoi(n)
    if err != nil {
        panic("ERROR")
    }
    name := ctx.PostForm("name")
    begintime := ctx.PostForm("begintime")
    finishtime := ctx.PostForm("finishtime")
    todo := ctx.PostForm("todo")
    models.DbUpdate(id, name, begintime, finishtime, todo)
    ctx.Redirect(302, "/")
}

//削除確認
func Deletecheck(ctx *gin.Context) {
    n := ctx.Param("id")
    id, err := strconv.Atoi(n)
    if err != nil {
        panic("ERROR")
    }
    himajin := models.DbGetOne(id)
    ctx.HTML(http.StatusOK, "delete.html", gin.H{"himajins": himajin})
}

//Delete
func Delete(ctx *gin.Context) {
    n := ctx.Param("id")
    id, err := strconv.Atoi(n)
    if err != nil {
        panic("ERROR")
    }
    models.DbDelete(id)
    ctx.Redirect(302, "/")
}
