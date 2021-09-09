#!/bin/bash

migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/bookings?sslmode=disable" -verbose up