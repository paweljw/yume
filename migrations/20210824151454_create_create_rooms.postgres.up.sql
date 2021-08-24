CREATE TABLE rooms (
    id integer PRIMARY KEY,
    description text,
    zone_id integer,
    CONSTRAINT fk_zone FOREIGN KEY(zone_id) REFERENCES zones(id)
)
