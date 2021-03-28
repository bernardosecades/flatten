#!/bin/bash
set -e
export PGPASSWORD=$POSTGRES_PASSWORD;
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname $POSTGRES_DB <<-EOSQL
  CREATE DATABASE flatten;
  GRANT ALL PRIVILEGES ON DATABASE flatten TO postgres;
  \connect flatten postgres
  BEGIN;
    CREATE TABLE IF NOT EXISTS history (
	  id CHAR(36) NOT NULL CHECK (CHAR_LENGTH(id) = 36) PRIMARY KEY,
    request bytea,
    response bytea,
    depth smallint,
    created_at timestamp
	);
	CREATE INDEX created_at_idx ON history(created_at);
  COMMIT;
EOSQL
