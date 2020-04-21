package models

import(
  "gopractice/crypto"
  "fmt"
  "errors"
  "gopractice/config"
  "github.com/jinzhu/gorm"
)

type User struct {
    gorm.Model
    Username string
    Email string
    Password string
    authenticated bool
}
func NewUser(username, email string) *User {
    return &User{
        Username: username,
        Email: email,
    }
}
func (u *User) SetPassword(password string) error {
    hash, err := crypto.PasswordEncrypt(password)
    if err != nil {
        return err
    }
    u.Password = hash
    return nil
}

func (u *User) Authenticate() {
    u.authenticated = true
}

////
//DB一つ取得
func UserDbGetOne(username, password string) (User, error) {
    db, err := gorm.Open(config.GetDBConfig())
    if err != nil {
        panic("データベース失敗(UserdbGetOne())")
    }
    var user User
    if result := db.Where("Username = ?", username).First(&user); result.Error != nil {
       fmt.Println("ユーザーが存在しません")
    }
    if err := crypto.CompareHashAndPassword(user.Password, password); err != nil {
        return user, errors.New("user \"" + username + "\" doesn't exists")
    }
    db.Close()
    return user, nil
}
//DBマイグレート
func UserDbInit() {
    db, err := gorm.Open(config.GetDBConfig())
    if err != nil {
        panic("データベース失敗（userdbInit）")
    }
    db.AutoMigrate(&User{})
    defer db.Close()
}
//DB追加
func UserDbInsert(username, email, password string) error{
    db, err := gorm.Open(config.GetDBConfig())
    if err != nil {
        panic("データベース失敗（UserdbInsert)")
    }
    db.Create(&User{Username: username, Email: email, Password: password})
    defer db.Close()
    return nil
}

// //DB更新
// func UserDbUpdate(id int, name string, begintime string, finishtime string, todo string) {
//     db, err := gorm.Open(config.GetDBConfig())
//     if err != nil {
//         panic("データベース失敗（dbUpdate)")
//     }
//     var himajin Himajin
//     db.First(&himajin, id)
//     himajin.Name = name
//     himajin.Begintime = begintime
//     himajin.Finishtime = finishtime
//     himajin.Todo = todo
//     db.Save(&himajin)
//     db.Close()
// }
//
// //DB削除
// func UserDbDelete(id int) {
//     db, err := gorm.Open(config.GetDBConfig())
//     if err != nil {
//         panic("データベース失敗（dbDelete)")
//     }
//     var himajin Himajin
//     db.First(&himajin, id)
//     db.Delete(&himajin)
//     db.Close()
// }
//
// //DB全取得
// func UserDbGetAll() []Himajin {
//     db, err := gorm.Open(config.GetDBConfig())
//     if err != nil {
//         panic("データベース失敗(dbGetAll())")
//     }
//     var himajins []Himajin
//     db.Order("created_at desc").Find(&himajins)
//     db.Close()
//     return himajins
// }

//DBからusername一致するもの一つ取得
func UserDbExists(username string) (User, bool) {
    db, err := gorm.Open(config.GetDBConfig())
    if err != nil {
        panic("データベース失敗(UserdbExists())")
    }
    var user User
    if result := db.First(&user, username); result.Error != nil {
       return user, false
    }
    db.Close()
    return user, true
}
