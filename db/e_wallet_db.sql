DROP DATABASE e_wallet_db;

CREATE DATABASE e_wallet_db;

SELECT * FROM users;
SELECT * FROM wallets;
SELECT * FROM transactions;
SELECT * FROM boxes;

SELECT * FROM forgot_password_tokens;
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


SELECT * FROM "transactions" 
WHERE "transactions"."deletd_at" IS NULL 
ORDER BY amount ASC 
LIMIT 25;

SELECT * FROM transactions;
SELECT * FROM boxes;

INSERT INTO boxes (amount)
VALUES
    (10000.00),
    (20000.00),
    (30000.00),
    (40000.00),
    (50000.00),
    (60000.00),
    (70000.00),
    (80000.00),
    (90000.00),
    (100000.00),
    (110000.00),
    (120000.00),
    (130000.00),
    (140000.00),
    (150000.00),
    (160000.00),
    (170000.00),
    (180000.00);