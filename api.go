package main

import (
	"fmt"
	"net/http"
  "time"
  "errors"
	//"github.com/go-martini/martini"

)

// RegisterEvent return the major id
func RegisterEvent(r *http.Request, enc Encoder) string {
  event, err := getPostEvent(r)
  fmt.Printf("%s\n",err)
  if err != nil {
    return Must(enc.Encode(NewError(400, fmt.Sprintf("%s", err))))
  }
  event.Major = 1111
  event.CreatedAt = time.Now()
  AddEvent(event)
	return Must(enc.Encode(event))
}

// Parse the request body, check input data
func getPostEvent(r *http.Request) (*Event, error) {
  en,rn,desc,items := r.FormValue("eventName"),r.FormValue("roomName"),r.FormValue("description"),r.FormValue("items")
  //var err error
  if en == "" {
    return nil, errors.New("Error: eventName is missing")
  }
  if items == "" {
    return nil, errors.New("Error: items is missing")
  }

  return &Event{
      EventName: en,
      RoomName:	rn,
      Description: desc,
      Items: items,
  }, nil
}

func checkString(s string) error {
  if s == "" {
    return errors.New("missing")
  }
  return nil
}
