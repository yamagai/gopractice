package models

import(
  "gopractice/config"
  "github.com/jinzhu/gorm"
)

type Post struct {
    gorm.Model
    PostUsername string
    Begintime string
    Finishtime string
    Todo string
}

//DBマイグレート
func DbInit() {
    db, err := gorm.Open(config.GetDBConfig())
    if err != nil {
        panic("データベース失敗（dbInit）")
    }
    db.AutoMigrate(&Post{})
    defer db.Close()
}
//DB追加
func DbInsert(username string, begintime string, finishtime string, todo string) {
    db, err := gorm.Open(config.GetDBConfig())
    if err != nil {
        panic("データベース失敗（dbInsert)")
    }
    db.Create(&Post{PostUsername: username, Begintime: begintime, Finishtime: finishtime, Todo: todo})
    defer db.Close()
}

//DB更新
func DbUpdate(id int, begintime string, finishtime string, todo string) {
    db, err := gorm.Open(config.GetDBConfig())
    if err != nil {
        panic("データベース失敗（dbUpdate)")
    }
    var post Post
    db.First(&post, id)
    post.Begintime = begintime
    post.Finishtime = finishtime
    post.Todo = todo
    db.Save(&post)
    db.Close()
}

//DB削除
func DbDelete(id int) {
    db, err := gorm.Open(config.GetDBConfig())
    if err != nil {
        panic("データベース失敗（dbDelete)")
    }
    var post Post
    db.First(&post, id)
    db.Delete(&post)
    db.Close()
}

//DB全取得
func DbGetAll() []Post {
    db, err := gorm.Open(config.GetDBConfig())
    if err != nil {
        panic("データベース失敗(dbGetAll())")
    }
    var posts []Post
    db.Order("created_at desc").Find(&posts)
    db.Close()
    return posts
}

//DB一つ取得
func DbGetOne(id int) Post {
    db, err := gorm.Open(config.GetDBConfig())
    if err != nil {
        panic("データベース失敗(dbGetOne())")
    }
    var post Post
    db.First(&post, id)
    db.Close()
    return post
}
