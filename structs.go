package structs

type Event struct {
  Id int64
  EventName string
  RoomName string
  Description string
  Items Item[]
  Major int16
  Date datetime
}

type User struct {
  Id int64
  UserName string
  Profile string

}

type Item struct {



}
