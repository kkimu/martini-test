package main

import(
    "time"
)

type Event struct {
  Id int64
  EventName string
  RoomName string
  Description string
  Items string
  Major int16
  CreatedAt time.Time
}

type User struct {
  Id int64
  UserName string
  Profile string
}

type Error struct {
  Message string
  Code int
}
