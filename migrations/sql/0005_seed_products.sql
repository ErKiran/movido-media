-- +goose Up
-- SQL in this section is executed when the migration is applied.
INSERT INTO products (product_code, product_name, price, currency)
VALUES
  ('PRD-141', 'VirtuoVision', 99.00, 'EUR'),
  ('PRD-145', 'NetNexus', 99.00, 'EUR'),
  ('PRD-160', 'CyberStream', 103.00, 'EUR');

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DELETE * FROM products;