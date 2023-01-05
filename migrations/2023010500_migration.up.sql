BEGIN;

CREATE TABLE basic_user (
    "id" SERIAL PRIMARY KEY,
    "username" VARCHAR NOT NULL,
    "password" VARCHAR NOT NULL
);

COMMIT;
