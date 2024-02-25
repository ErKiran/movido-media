-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE IF NOT EXISTS products(
    product_code varchar(16) PRIMARY KEY,
    product_name varchar(64) NOT NULL, 
    price int NOT NULL,
    currency varchar(16) NOT NULL,
    created_date timestamptz NOT NULL DEFAULT now(),
	modified_date timestamptz NOT NULL DEFAULT now()
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE products;