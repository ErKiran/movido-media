-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS customers(
    customer_id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    full_name varchar(64) NOT NULL, 
    email varchar(64) NOT NULL,
    address varchar(64) NOT NULL,
    created_date timestamptz NOT NULL DEFAULT now(),
	modified_date timestamptz NOT NULL DEFAULT now()
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE customers;