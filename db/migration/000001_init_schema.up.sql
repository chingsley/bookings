CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "email" VARCHAR (300) UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "access_level" bigint NOT NULL DEFAULT (1),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "rooms" (
  "id" bigserial PRIMARY KEY,
  "room_name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "restrictions" (
  "id" bigserial PRIMARY KEY,
  "restrictions_name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "reservations" (
  "id" bigserial PRIMARY KEY,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "email" VARCHAR (300) UNIQUE NOT NULL,
  "phone" varchar NOT NULL,
  "start_date" timestamptz NOT NULL,
  "end_date" timestamptz NOT NULL,
  "room_id" bigint NOT NULL DEFAULT (1),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "room_restrictions" (
  "id" bigserial PRIMARY KEY,
  "start_date" timestamptz NOT NULL,
  "end_date" timestamptz NOT NULL,
  "room_id" bigint NOT NULL DEFAULT (1),
  "reservation_id" bigint NOT NULL DEFAULT (1),
  "restriction_id" bigint NOT NULL DEFAULT (1),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "reservations" ADD FOREIGN KEY ("room_id") REFERENCES "rooms" ("id");
ALTER TABLE "room_restrictions" ADD FOREIGN KEY ("room_id") REFERENCES "rooms" ("id");
ALTER TABLE "room_restrictions" ADD FOREIGN KEY ("reservation_id") REFERENCES "reservations" ("id");
ALTER TABLE "room_restrictions" ADD FOREIGN KEY ("restriction_id") REFERENCES "restrictions" ("id");


CREATE INDEX ON "room_restrictions" ("start_date");
CREATE INDEX ON "room_restrictions" ("end_date");
CREATE INDEX ON "room_restrictions" ("room_id");
CREATE INDEX ON "room_restrictions" ("reservation_id");
CREATE INDEX ON "reservations" ("email");
CREATE INDEX ON "reservations" ("last_name");