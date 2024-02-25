-- +goose Up
-- SQL in this section is executed when the migration is applied.
INSERT INTO customers (customer_id, full_name, email, address)
VALUES
  ('123e4567-e89b-12d3-a456-426614174000', 'John Doe', 'john.doe@example.com', '123 Main St'),
  ('223e4567-e89b-12d3-a456-426614174001', 'Jane Smith', 'jane.smith@example.com', '456 Elm St'),
  ('323e4567-e89b-12d3-a456-426614174002', 'Alice Johnson', 'alice.johnson@example.com', '789 Oak St');

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DELETE * FROM customers;