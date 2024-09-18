BEGIN;

BEGIN;
CREATE TABLE IF NOT EXISTS product
(
    id         SERIAL PRIMARY KEY NOT NULL,
    name       VARCHAR NOT NULL UNIQUE,
    cost       INTEGER NOT NULL,
    amount     INTEGER NOT NULL,
    updated_at DATE DEFAULT NULL,
    created_at DATE NOT NULL
);
COMMIT;

BEGIN;
insert into product
    (name, cost, amount, updated_at, created_at)
values
    ('А тебя ебать не должно как это называется', 1200000, 1, null, CURRENT_DATE),
    ('А тебя ебать не должно как это называется x 2', 100, 100, null, CURRENT_DATE),
    ('А тебя ебать не должно как это называется x 3', 150, 1, null, CURRENT_DATE);
COMMIT;

COMMIT;