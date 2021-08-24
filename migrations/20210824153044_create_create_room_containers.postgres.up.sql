CREATE TABLE room_containers (
  id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  container_id integer NOT NULL,
  room_id integer NOT NULL,
  constraint fk_container foreign key(container_id) references containers(id),
  constraint fk_room foreign key(room_id) references rooms(id)
)
