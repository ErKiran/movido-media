-- +goose Up
-- SQL in this section is executed when the migration is applied.
INSERT INTO contracts (contract_id, product_code, customer_id, start_date, duration, billing_frequency, billing_frequency_units, duration_units)
VALUES
  ('123e4567-e89b-12d3-a456-426614174006', 'PRD-141', '123e4567-e89b-12d3-a456-426614174000', '2025-02-01', 12, 3, 'MONTHS', 'MONTHS'),
  ('223e4567-e89b-12d3-a456-426614174007', 'PRD-145', '223e4567-e89b-12d3-a456-426614174001', '2024-01-15', 24, 1, 'MONTHS', 'MONTHS'),
  ('323e4567-e89b-12d3-a456-426614174008', 'PRD-160', '323e4567-e89b-12d3-a456-426614174002', '2024-02-26', 36, 3, 'MONTHS', 'MONTHS');


-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DELETE * FROM contracts;