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