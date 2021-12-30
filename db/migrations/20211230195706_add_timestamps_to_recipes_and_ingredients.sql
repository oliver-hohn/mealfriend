-- +goose Up
-- +goose StatementBegin
ALTER TABLE recipes
  ADD created_at TIMESTAMP WITH TIME ZONE NOT NULL,
  ADD updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
  ADD deleted_at TIMESTAMP WITH TIME ZONE;
CREATE INDEX index_recipes_on_deleted_at ON recipes(deleted_at);

ALTER TABLE ingredients
  ADD created_at TIMESTAMP WITH TIME ZONE NOT NULL,
  ADD updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
  ADD deleted_at TIMESTAMP WITH TIME ZONE;
CREATE INDEX index_ingredients_on_deleted_at ON ingredients(deleted_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX index_ingredients_on_deleted_at;
ALTER TABLE ingredients
  DROP created_at,
  DROP updated_at,
  DROP deleted_at;

DROP INDEX index_recipes_on_deleted_at;
ALTER TABLE recipes
  DROP created_at,
  DROP updated_at,
  DROP deleted_at;
-- +goose StatementEnd
