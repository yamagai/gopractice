package routes

import (
    "gopractice/models"
    "gopractice/sessions"
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)

//ログイン
func UserLogIn(ctx *gin.Context) {
    println("post/login")
    username := ctx.PostForm("username")
    password := ctx.PostForm("password")

    user, err := models.UserDbGetOne(username, password)
    if err != nil {
        println("Error: " + err.Error())
        ctx.Redirect(http.StatusSeeOther, "/")
        return
    }

    println("Authentication Success!!")
    println("  username: " + user.Username)
    println("  email: " + user.Email)
    println("  password: " + user.Password)
    session := sessions.GetDefaultSession(ctx)
    session.Set("user", &user)
    session.Save()
    user.Authenticate()

    ctx.Redirect(302, "/")
}

//ユーザー登録
func UserSignUp(ctx *gin.Context) {
    username := ctx.PostForm("username")
    email := ctx.PostForm("emailaddress")
    password := ctx.PostForm("password")
    passwordConf := ctx.PostForm("passwordconfirmation")
    if password != passwordConf {
        println("Error: password and passwordConf not match")
        ctx.Redirect(http.StatusSeeOther, "/")
        return
    }
    alreadyuser, exists := models.UserDbExists(username)
    if exists{
        fmt.Println("ユーザー名 \"" + alreadyuser.Username + "\" はすでに使用されています")
        return
    }
    user := models.NewUser(username, email)
    if err := user.SetPassword(password); err != nil {
        return
    }

    if err := models.UserDbInsert(user.Username, user.Email, user.Password); err != nil {
        println("Error: " + err.Error())
    } else {
        println("Signup success!!")
        println("  username: " + user.Username)
        println("  email: " + user.Email)
        println("  password: " + user.Password)
    }
    session := sessions.GetDefaultSession(ctx)
    session.Set("user", user)
    session.Save()
    println("Session saved.")
    println("  sessionID: " + session.ID)
    ctx.Redirect(302, "/")
}

//ログアウト
func UserLogOut(ctx *gin.Context) {
    session := sessions.GetDefaultSession(ctx)
    session.Terminate()
    ctx.Redirect(http.StatusSeeOther, "/")
}
