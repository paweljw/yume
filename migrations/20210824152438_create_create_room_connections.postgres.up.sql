CREATE TABLE room_connections (
    id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    from_id integer,
    to_id integer,
    direction smallint,
    locked_by_id integer default null,
    locked_by_flag varchar(100),
    constraint fk_from foreign key(from_id) references rooms(id),
    constraint fk_to foreign key(to_id) references rooms(id),
    constraint fk_locked_by foreign key(locked_by_id) references items(id)
);
