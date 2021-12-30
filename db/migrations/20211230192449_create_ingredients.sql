-- +goose Up
-- +goose StatementBegin
CREATE TABLE ingredients(
  id bigint NOT NULL PRIMARY KEY,
  name text NOT NULL,
  type text,
  recipe_id bigint NOT NULL REFERENCES recipes(id)
);
CREATE INDEX index_ingredients_on_recipe_id ON recipes(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX index_ingredients_on_recipe_id;
DROP TABLE ingredients;
-- +goose StatementEnd
