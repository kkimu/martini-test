package main

import (
  "fmt"
  "net/http"
  "time"
  "gopkg.in/validator.v2"
)

// RegisterEvent return the major id
func RegisterEvent(r *http.Request, enc Encoder) string {
  event := getPostEvent(r)
  if errs := validator.Validate(event); errs != nil {
    fmt.Printf("errs: %s\n",errs)
    return Must(enc.Encode(NewError(400, fmt.Sprintf("%s",errs))))
  }

  //event.Major = 1111

  AddEvent(event)
	return Must(enc.Encode(event))
}

// Parse the request body, check input data
func getPostEvent(r *http.Request) *Event {
  // リクエストボディをパースして代入
  en,rn,desc,items := r.FormValue("eventName"),r.FormValue("roomName"),r.FormValue("description"),r.FormValue("items")

  return &Event{
      EventName: en,
      RoomName:	rn,
      Description: desc,
      Items: items,
      Active: true,
      CreatedAt: time.Now(),
  }
}
