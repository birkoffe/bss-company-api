-- +goose Up

INSERT INTO COMPANY (ID, NAME, ADDRESS, REMOVED, CREATED, UPDATED) VALUES
  (1, 'COMPANY_1', 'UNKNOWN_1', FALSE, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP), 
  (2, 'COMPANY_2', 'UNKNOWN_2', TRUE, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- +goose Down
