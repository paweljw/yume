CREATE TABLE container_inventories (
    id integer PRIMARY KEY,
    container_id integer,
    item_id integer,
    rate real default 1.0,
    constraint fk_container FOREIGN KEY(container_id) REFERENCES containers(id),
    constraint fk_item FOREIGN KEY(item_id) REFERENCES items(id)
)
