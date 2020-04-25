package routes

import (
  "gopractice/sessions"
  "gopractice/models"
  "net/http"
  "strconv"
  "github.com/gin-gonic/gin"
)

func LogIn(ctx *gin.Context) {
    ctx.HTML(http.StatusOK, "login.html", gin.H{})
}

func SignUp(ctx *gin.Context) {
    ctx.HTML(http.StatusOK, "signup.html", gin.H{})
}

func NoRoute(ctx *gin.Context) {
    ctx.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
}
//index
func Index(ctx *gin.Context) {
  post := models.DbGetAll()
  session := sessions.GetDefaultSession(ctx)
   buffer, exists := session.Get("user")
   if !exists {
        println("Unhappy home")
        println("  sessionID: " + session.ID)
        session.Save()
        ctx.HTML(http.StatusOK, "index.html", gin.H{
          "Posts": post,
        })
        return
    }
   user := buffer.(*models.User)
   println("Home sweet home")
   println("  sessionID: " + session.ID)
   println("  username: " + user.Username)
   println("  email: " + user.Email)
   session.Save()
    ctx.HTML(http.StatusOK, "index.html", gin.H{
        "Posts": post,
        "isLoggedIn": exists,
        "username": user.Username,
        "email": user.Email,
    })
}

//create
func Create(ctx *gin.Context) {
    username := ctx.PostForm("username")
    begintime := ctx.PostForm("begintime")
    finishtime := ctx.PostForm("finishtime")
    todo := ctx.PostForm("todo")
    models.DbInsert(username, begintime, finishtime, todo)
    ctx.Redirect(302, "/")
}

//Detail
func Detail(ctx *gin.Context) {
    n := ctx.Param("id")
    id, err := strconv.Atoi(n)
    if err != nil {
        panic(err)
    }
    post := models.DbGetOne(id)
    ctx.HTML(http.StatusOK, "detail.html", gin.H{"Posts": post})
}

//update
func Update(ctx *gin.Context) {
    n := ctx.Param("id")
    id, err := strconv.Atoi(n)
    if err != nil {
        panic("ERROR")
    }
    begintime := ctx.PostForm("begintime")
    finishtime := ctx.PostForm("finishtime")
    todo := ctx.PostForm("todo")
    models.DbUpdate(id, begintime, finishtime, todo)
    ctx.Redirect(302, "/")
}

//削除確認
func Deletecheck(ctx *gin.Context) {
    n := ctx.Param("id")
    id, err := strconv.Atoi(n)
    if err != nil {
        panic("ERROR")
    }
    Post := models.DbGetOne(id)
    ctx.HTML(http.StatusOK, "delete.html", gin.H{"Posts": Post})
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
