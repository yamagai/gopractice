package routes

import (
    "gopractice/models"
    "gopractice/config"
    "errors"
    "net/http"
    "github.com/gin-gonic/gin"
)

//ログイン
func UserLogIn(ctx *gin.Context) {
    println("post/login")
    username := ctx.PostForm("username")
    password := ctx.PostForm("password")

    user, err := models.UserdbGetOne(username, password)
    if err != nil {
        println("怪しいとこや: " + err.Error())
    } else {
        println("Authentication Success!!")
        println("  username: " + user.Username)
        println("  email: " + user.Email)
        println("  password: " + user.Password)
        user.Authenticate()
    }

    ctx.Redirect(302, "/")
}

//create(ユーザー登録)
func UserSignUp(ctx *gin.Context) {
    username := ctx.PostForm("username")
    email := ctx.PostForm("emailaddress")
    password := ctx.PostForm("password")
    passwordConf := ctx.PostForm("passwordconfirmation")
    if password != passwordConf {
        println("Error: password and passwordConf not match")
        ctx.Redirect(http.StatusSeeOther, "//localhost:8080/")
        return
    }
    if models.UserDbExists(username) {
        return errors.New("ユーザー名 \"" + username + "\" はすでに使用されています")
    }
    user := models.NewUser(username, email)
    if err := user.SetPassword(password); err != nil {
        return err
    }

    if err := models.UserDbInsert(user.Username, user.Email, user.Password); err != nil {
        println("Error: " + err.Error())
    } else {
        println("Signup success!!")
        println("  username: " + user.Username)
        println("  email: " + user.Email)
        println("  password: " + user.Password)
    }
    ctx.Redirect(302, "/")
}

// //Detail
// func UserDetail(ctx *gin.Context) {
//     n := ctx.Param("id")
//     id, err := strconv.Atoi(n)
//     if err != nil {
//         panic(err)
//     }
//     himajin := models.DbGetOne(id)
//     ctx.HTML(http.StatusOK, "detail.html", gin.H{"himajins": himajin})
// }
//
// //update
// func UserUpdate(ctx *gin.Context) {
//     n := ctx.Param("id")
//     id, err := strconv.Atoi(n)
//     if err != nil {
//         panic("ERROR")
//     }
//     name := ctx.PostForm("name")
//     begintime := ctx.PostForm("begintime")
//     finishtime := ctx.PostForm("finishtime")
//     todo := ctx.PostForm("todo")
//     models.DbUpdate(id, name, begintime, finishtime, todo)
//     ctx.Redirect(302, "/")
// }
//
// //削除確認
// func UserDeletecheck(ctx *gin.Context) {
//     n := ctx.Param("id")
//     id, err := strconv.Atoi(n)
//     if err != nil {
//         panic("ERROR")
//     }
//     himajin := models.DbGetOne(id)
//     ctx.HTML(http.StatusOK, "delete.html", gin.H{"himajins": himajin})
// }
//
// //Delete
// func UserDelete(ctx *gin.Context) {
//     n := ctx.Param("id")
//     id, err := strconv.Atoi(n)
//     if err != nil {
//         panic("ERROR")
//     }
//     models.DbDelete(id)
//     ctx.Redirect(302, "/")
// }
