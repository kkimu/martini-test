package main

import (
  "github.com/go-martini/martini"
  _ "github.com/go-sql-driver/mysql"
  "github.com/jinzhu/gorm"
  "github.com/martini-contrib/render"
)

type Item struct {
  Id int64
  Title string
  Description string
  UserName string
}

var (
  db gorm.DB
  sqlConnection string
)

func main() {
  var err error

  db, err = gorm.Open("mysql", "root:pass@tcp(127.0.0.1:3306)/airmeet?parseTime=True")

  if err != nil {
    panic(err)
    return
  }

  m := martini.Classic()

  m.Use(render.Renderer())

  m.Get("/", func(r render.Render) {
    var retData struct {
      Items []Item
    }

    db.Find(&retData.Items)

    r.JSON(200, map[string]interface{}{"hello": "world!"})
  })
  m.Run()
}
