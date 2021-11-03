-- +goose Up
CREATE TABLE company (
  id BIGSERIAL PRIMARY KEY,
  foo BIGINT NOT NULL
);

-- +goose Down
DROP TABLE company;
