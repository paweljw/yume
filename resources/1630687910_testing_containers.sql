INSERT INTO containers VALUES (
  1,
  'black box',
  'It is very black.',
  false
);

INSERT INTO containers VALUES (
  2,
  'Floor for room 2',
  'Floor for room 2',
  false
);

INSERT INTO room_containers (container_id, room_id)
VALUES (1, 2);

INSERT INTO room_containers (container_id, room_id)
VALUES (2, 2);

INSERT INTO items VALUES (
  1,
  'small bronze key',
  'A small bronze key. Looks like it opens something of not much import.',
  1,
  0
);

INSERT INTO items VALUES (
  2,
  'uncannily large seashell',
  'Rare loot. Only about 80% of people find it!',
  50,
  0
);

INSERT INTO container_inventories (container_id, item_id, rate)
VALUES (1, 1, 1.0);

INSERT INTO container_inventories (container_id, item_id, rate)
VALUES (2, 2, 0.8);
