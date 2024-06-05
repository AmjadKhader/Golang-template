-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied.

CREATE TABLE IF NOT EXISTS books (
    id              SERIAL PRIMARY KEY,
    title           VARCHAR(255) NOT NULL,
    auther          VARCHAR(255) NOT NULL,
    total_pages     INT NOT NULL
);

-- +goose Down
-- SQL in section 'Down' is executed when this migration is rolled back.

DROP TABLE IF EXISTS books;
