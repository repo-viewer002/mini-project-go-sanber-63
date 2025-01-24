-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE users (
  id          SERIAL PRIMARY KEY,
  first_name  varchar(255) NOT NULL,
  last_name   varchar(255) NOT NULL
)

-- +migrate StatementEnd