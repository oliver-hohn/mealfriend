-- +goose Up
-- +goose StatementBegin
CREATE TABLE recipes(
  id bigint NOT NULL PRIMARY KEY,
  code text NOT NULL,
  name text NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE recipes;
-- +goose StatementEnd
