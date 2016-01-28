package main

import(
    "time"
)

type Event struct {
  Id int64
  EventName string `validate:"nonzero"`
  RoomName string
  Description string
  Items string `validate:"nonzero"`
  Major int16
  CreatedAt time.Time
  DeletedAt *time.Time
}

type User struct {
  Id int64
  UserName string
  Profile string
}
