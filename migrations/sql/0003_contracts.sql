-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS contracts(
    contract_id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    product_code varchar(16) NOT NULL,
    customer_id UUID NOT NULL,
    start_date timestamptz NOT NULL,
    duration int NOT NULL,
    billing_frequency int NOT NULL,
    billing_frequency_units varchar(16) NOT NULL,
    duration_units varchar(16) NOT NULL,
    created_date timestamptz NOT NULL DEFAULT now(),
	modified_date timestamptz NOT NULL DEFAULT now(),
    FOREIGN KEY (product_code) REFERENCES products(product_code),
    FOREIGN KEY (customer_id) REFERENCES customers(customer_id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE contracts;