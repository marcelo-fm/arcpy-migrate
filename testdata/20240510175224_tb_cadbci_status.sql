-- +goose Up
CREATE TABLE cadbci_status (
  objectid serial PRIMARY KEY,
  codigo_cadastro integer NOT NULL UNIQUE,
  last_operation varchar,
  source varchar,
  status boolean,
  globalid varchar NOT NULL UNIQUE,
  created_date timestamp NOT NULL,
  last_edited_date timestamp NOT NULL
);

-- +goose Down
DROP TABLE cadbci_status;
