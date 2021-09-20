#!/bin/bash


psql bookings \
    -U postgres \
    -p 5432 \
    -h 127.0.0.1 \
    -c "\copy \"restrictions\" from './db/seed/restrictions.csv' with DELIMITER ',' CSV HEADER"
psql bookings \
    -U postgres \
    -p 5432 \
    -h 127.0.0.1 \
    -c "\copy \"rooms\" from './db/seed/rooms.csv' with DELIMITER ',' CSV HEADER"
psql bookings \
    -U postgres \
    -p 5432 \
    -h 127.0.0.1 \
    -c "\copy \"reservations\" from './db/seed/reservations.csv' with DELIMITER ',' CSV HEADER"
psql bookings \
    -U postgres \
    -p 5432 \
    -h 127.0.0.1 \
    -c "\copy \"room_restrictions\" from './db/seed/room_restrictions.csv' with DELIMITER ',' CSV HEADER"

psql bookings \
    -U postgres \
    -p 5432 \
    -h 127.0.0.1 \
    -c "SELECT setval('restrictions_id_seq', COALESCE((SELECT MAX(id)+1 FROM restrictions), 1), false)"

psql bookings \
    -U postgres \
    -p 5432 \
    -h 127.0.0.1 \
    -c "SELECT setval('rooms_id_seq', COALESCE((SELECT MAX(id)+1 FROM rooms), 1), false)"

psql bookings \
    -U postgres \
    -p 5432 \
    -h 127.0.0.1 \
    -c "SELECT setval('reservations_id_seq', COALESCE((SELECT MAX(id)+1 FROM reservations), 1), false)"

psql bookings \
    -U postgres \
    -p 5432 \
    -h 127.0.0.1 \
    -c "SELECT setval('room_restrictions_id_seq', COALESCE((SELECT MAX(id)+1 FROM room_restrictions), 1), false)"