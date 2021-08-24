CREATE TABLE room_current_inventories (
    id integer primary key generated always as identity,
    room_container_id integer not null,
    item_id integer not null,
    visible_to_id integer not null,
    constraint fk_room_container foreign key(room_container_id) references room_containers(id),
    constraint fk_item foreign key(item_id) references items(id),
    constraint fk_visible_to foreign key(visible_to_id) references players(id)
)
