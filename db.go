package main

import (
  //"fmt"
  _ "github.com/go-sql-driver/mysql"
  "github.com/jinzhu/gorm"
)

var db gorm.DB

func init() {
  var err error
  db, err = gorm.Open("mysql", "root:pass@tcp(127.0.0.1:3306)/airmeet?parseTime=True&loc=Japan")

  if err != nil {
    panic(err)
    return
  }
  db.DB()
}

func AddEvent(event *Event) {
  //db.NewRecord(event)
  db.Create(&event)
  //db.Save(&event)
}
