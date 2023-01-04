BEGIN;

CREATE TABLE basic_user (
    "id" SERIAL PRIMARY KEY,
    "user_name" VARCHAR NOT NULL
);

COMMIT;
