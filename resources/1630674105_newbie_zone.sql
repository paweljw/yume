INSERT INTO zones VALUES (
  1,
  'Newbie zone',
  'The zone for newbies. Some rooms or whatever.',
  true
);

INSERT INTO rooms VALUES (
  1,
  'It''s dark here. Rough-hewn stone walls glisten with moisture. The smell of rot and decay fills the air.\n\nIt helps with newline testing vehemently.',
  1
);

INSERT INTO rooms VALUES (
  2,
  'It''s also dark here. What, you expected proper descriptions from testing rooms? Please.',
  1
);

INSERT INTO room_connections (from_id, to_id, direction) VALUES (
  1,
  2,
  3 -- WEST
);

INSERT INTO room_connections (from_id, to_id, direction) VALUES (
  2,
  1,
  2 -- ACTUALLY FUCKING SOUTH, LEARN TO READ
);
