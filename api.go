package main

import (
	"fmt"
	"net/http"
	//"github.com/go-martini/martini"
)

// GetAlbums returns the list of albums (possibly filtered).
func RegisterEvent(r *http.Request, enc Encoder) string {
  event := Event{
    EventName:	r.FormValue("eventName"),
    RoomName:	r.FormValue("roomName"),
    Description: r.FormValue("description"),
    Items: r.FormValue("items"),
  }
	return Must(enc.Encode(event))
}
