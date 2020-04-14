package main

import (
  "time"
  "net/http"
  "github.com/gin-gonic/gin"
)

type ToDo struct {
  Hito string
  Content string
  Status int
  CreatedAt time.Time
}

func main() {
  todo := []ToDo{
    ToDo {
      Hito: "yamamoto",
      Content: "早起き",
      Status: 0,
      CreatedAt: time.Now(),
    },
    ToDo {
      Hito: "katsuhei",
      Content: "インターン応募",
      Status: 0,
      CreatedAt: time.Now(),
    },
    ToDo {
      Hito: "kohei",
      Content: "案件とる",
      Status: 0,
      CreatedAt: time.Now(),
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
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
