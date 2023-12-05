DROP DATABASE e_wallet_db;

CREATE DATABASE e_wallet_db;

SELECT * FROM users;
SELECT * FROM wallets;
SELECT * FROM transactions;

UPDATE wallets
SET balance = 50000.00;

SELECT * FROM source_of_funds;

INSERT INTO source_of_funds (source)
VALUES
    ('Bank Transfer'),
    ('Credit Card'),
    ('Cash'),
    ('Reward'),
    ('Other');