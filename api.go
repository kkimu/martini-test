package main

import (
	//"fmt"
	"net/http"
  "time"
	//"github.com/go-martini/martini"

)

// GetAlbums returns the list of albums (possibly filtered).
func RegisterEvent(r *http.Request, enc Encoder) string {
  event, err := getParams(r)
  if err != nil {
    return Must(enc.Encode(Error{
        Message: err,
        Code: 400,
      }))
  }
  event.Major = 1111
  event.CreatedAt = time.Now()
  AddEvent(event)
	return Must(enc.Encode(event))
}

func getParams(r *http.Request) (*Event, error) {
  en,rn,desc,items := r.FormValue("eventName"),r.FormValue("roomName"),r.FormValue("description"),r.FormValue("items")
  var err error
  if en == nil {
    err := "Error: eventName is missing"
  }
  if rm == nil {
    err := "Error: roomName is missing"
  }
  if err != nil {
    return nil, err
  } else {
    return &Event{
      EventName: en,
      RoomName:	rn,
      Description: desc,
      Items: items,
    }, nil
  }
}
