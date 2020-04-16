package models

import(
  "github.com/jinzhu/gorm"
)

type Himajin struct {
    gorm.Model
    Name   string
    Begintime string
    Finishtime string
    Todo string
}

//DBマイグレート
func DbInit() {
    db, err := gorm.Open("mysql", "root:@(localhost)/gopractice?charset=utf8&parseTime=True")
    if err != nil {
        panic("データベース失敗（dbInit）")
    }
    db.AutoMigrate(&Himajin{})
    defer db.Close()
}
//DB追加
func DbInsert(name string, begintime string, finishtime string, todo string) {
    db, err := gorm.Open("mysql", "root:@(localhost)/gopractice?charset=utf8&parseTime=True")
    if err != nil {
        panic("データベース失敗（dbInsert)")
    }
    db.Create(&Himajin{Name: name, Begintime: begintime, Finishtime: finishtime, Todo: todo})
    defer db.Close()
}

//DB更新
func DbUpdate(id int, name string, begintime string, finishtime string, todo string) {
    db, err := gorm.Open("mysql", "root:@(localhost)/gopractice?charset=utf8&parseTime=True")
    if err != nil {
        panic("データベース失敗（dbUpdate)")
    }
    var himajin Himajin
    db.First(&himajin, id)
    himajin.Name = name
    himajin.Begintime = begintime
    himajin.Finishtime = finishtime
    himajin.Todo = todo
    db.Save(&himajin)
    db.Close()
}

//DB削除
func DbDelete(id int) {
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
func DbGetAll() []Himajin {
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
func DbGetOne(id int) Himajin {
    db, err := gorm.Open("mysql", "root:@(localhost)/gopractice?charset=utf8&parseTime=True")
    if err != nil {
        panic("データベース失敗(dbGetOne())")
    }
    var himajin Himajin
    db.First(&himajin, id)
    db.Close()
    return himajin
}
