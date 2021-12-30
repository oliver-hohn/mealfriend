-- +goose Up
-- +goose StatementBegin
-- Create a sequence for the PK:
-- https://stackoverflow.com/questions/31965506/postgresql-column-type-conversion-from-bigint-to-bigserial
CREATE SEQUENCE recipes_id_seq;
ALTER TABLE recipes ALTER id SET DEFAULT nextval('recipes_id_seq');
ALTER SEQUENCE recipes_id_seq OWNED BY recipes.id;

CREATE SEQUENCE ingredients_id_seq;
ALTER TABLE ingredients ALTER id SET DEFAULT nextval('ingredients_id_seq');
ALTER SEQUENCE ingredients_id_seq OWNED BY ingredients.id;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE ingredients ALTER id SET DEFAULT null;
DROP SEQUENCE ingredients_id_seq;

ALTER TABLE recipes ALTER id SET DEFAULT null;
DROP SEQUENCE recipes_id_seq;
-- +goose StatementEnd
