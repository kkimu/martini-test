package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/jinzhu/gorm"
	"github.com/go-sql-driver/mysql"
  "strings"
  "fmt"
)

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())

	m.Get("/", func(r render.Render) {
		r.JSON(200, map[string]interface{}{"hello": "world!"})
	})

	m.Run()
}
