package main

import (
  "fmt"
  "time"
  "net/http"
  "github.com/gin-gonic/gin"
  _ "github.com/go-sql-driver/mysql"
  . "github.com/volatiletech/sqlboiler/queries/qm"
  "github.com/volatiletech/sqlboiler/boil"

)

type ToDo struct {
  ID int `form:"id"`
  Hito string `form:"hito"`
  Content string `form:"content"`
  Status int `form:"status"`
  CreatedAt time.Time
  CreatedAtS string
}

var todo []ToDo
var idMax = 0

func Saiban() int{
  idMax = idMax + 1
  return idMax
}
func GetDataToDo(c *gin.Context) {
    var b ToDo
    c.Bind(&b)
    b.ID = Saiban()
    b.Status = 0
    b.CreatedAtS = time.Now().Format("2006-01-02 15:04:05")
    todo = append(todo, b)
    c.HTML(http.StatusOK, "index.html", map[string]interface{}{
      "todo": todo,
    })
}
func GetDoneToDo(c *gin.Context) {
  var b ToDo
  if err := c.Bind(&b); err != nil {
    fmt.Errorf("%#v", err)
  }
  var s int

  if b.Status == 0 {
    s = 1
  } else {
    s = 0
  }

  for idx, t := range todo {
    if t.ID == b.ID {
      todo[idx].Status = s
    }
  }

    c.HTML(http.StatusOK, "index.html", map[string]interface{}{
      "todo": todo,
    })
}
func main() {
  todo = []ToDo{
    ToDo {
      ID: Saiban(),
      Hito: "yamamoto",
      Content: "早起き",
      Status: 0,
      CreatedAt: time.Now(),
      CreatedAtS: time.Now().Format("2006-01-02 15:04:05"),
    },
    ToDo {
      ID: Saiban(),
      Hito: "katsuhei",
      Content: "インターン応募",
      Status: 1,
      CreatedAt: time.Now(),
      CreatedAtS: time.Now().Format("2006-01-02 15:04:05"),
    },
    ToDo {
      ID: Saiban(),
      Hito: "kohei",
      Content: "案件とる",
      Status: 0,
      CreatedAt: time.Now(),
      CreatedAtS: time.Now().Format("2006-01-02 15:04:05"),
    },
  }


  r := gin.Default()
  r.LoadHTMLFiles("./templates/index.html")

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
  r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", map[string]interface{}{
      "todo": todo,
    })
	})
  r.GET("/yaru",GetDataToDo)
  r.GET("/done",GetDoneToDo)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
