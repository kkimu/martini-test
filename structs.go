package main

import(
    "time"
)

type Event struct {
  Id int64 `json:"id"`
  EventName string `json:"event_name" validate:"nonzero"`
  RoomName string `json:"room_name"`
  Description string `json:"description"`
  Items string `json:"items" validate:"nonzero"`
  Major int16 `json:"major"`
  Active bool `json:"active" validate:"nonzero"`
  CreatedAt time.Time `json:"created_at"`
  DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type User struct {
  Id int64
  UserName string
  Profile string
}
