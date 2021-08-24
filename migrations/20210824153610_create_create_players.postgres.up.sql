CREATE TABLE players (
    id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name varchar(20) NOT NULL,
    password char(64) NOT NULL,
    race smallint DEFAULT 0,
    pronouns smallint DEFAULT 0,
    saved_room_id integer,
    current_room_id integer,
    constraint fk_saved_room foreign key(saved_room_id) references rooms(id),
    constraint fk_current_room foreign key(current_room_id) references rooms(id)
)
