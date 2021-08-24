CREATE TABLE player_current_inventories (
    id integer primary key generated always as identity,
    player_id integer not null,
    item_id integer not null,
    is_bound boolean default false,
    is_equipped boolean default false,
    constraint fk_player foreign key(player_id) references players(id),
    constraint fk_item foreign key(item_id) references items(id)
)
