drop database airmeet;
create database airmeet;
use airmeet;

create table events(
  id int(10) auto_increment primary key,
  event_name varchar(64) not null,
  room_name varchar(64),
  description varchar(64),
  items varchar(64),
  major int(5),
  created_at datetime
)
