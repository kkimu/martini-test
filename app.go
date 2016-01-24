package main

import (
  "github.com/go-martini/martini"
  _ "github.com/go-sql-driver/mysql"
  "github.com/jinzhu/gorm"
  "github.com/martini-contrib/render"
  "fmt"
	"net/http"
)

type Event struct {
  Id int64
  EventName string
  RoomName string
  Description string
  Items string
  Major int16
  Date string
}

type User struct {
  Id int64
  UserName string
  Profile string

}

func main() {
  db, err := gorm.Open("mysql", "root:pass@tcp(127.0.0.1:3306)/airmeet?parseTime=True")

  if err != nil {
    panic(err)
    return
  }

	db.DB()
  m := martini.Classic()

  m.Use(render.Renderer())

  m.Get("/", func(r render.Render) {
    r.JSON(200, map[string]interface{}{"hello": "world!"})
  })

	m.Post("/registerEvent", func(req *http.Request,r render.Render) {
		event := Event{
	    EventName:	req.FormValue("eventName"),
			RoomName:	req.FormValue("roomName"),
			Description: req.FormValue("description"),
			Items: req.FormValue("items"),
	  }
		fmt.Printf("a")
		r.JSON(200, event)
	})
  m.Run()
}
