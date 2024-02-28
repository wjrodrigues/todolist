#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
	CREATE DATABASE todolist_test;
	GRANT ALL PRIVILEGES ON DATABASE todolist_test TO simple_user;
EOSQL
