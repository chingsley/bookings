#!/bin/bash

migrate -path db/migration -database "postgresql://kingsleyeneja:chinonxo@localhost:5432/bookings?sslmode=disable" -verbose down