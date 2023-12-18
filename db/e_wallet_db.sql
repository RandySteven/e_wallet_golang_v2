DROP DATABASE e_wallet_db;

CREATE DATABASE e_wallet_db;

SELECT * FROM users;
SELECT * FROM wallets;
SELECT * FROM transactions;
SELECT * FROM boxes;

SELECT * FROM games WHERE user_id = 11;
SELECT * FROM users WHERE id = 11;
SELECT * FROM forgot_password_tokens;
SELECT * FROM wallets WHERE id = 11;

UPDATE wallets
SET balance = 0
WHERE id = 11;

UPDATE users
SET chance = 0
WHERE id = 11;

SELECT * FROM source_of_funds;

SELECT * FROM wallets;

INSERT INTO source_of_funds (source)
VALUES
    ('Bank Transfer'),
    ('Credit Card'),
    ('Cash'),
    ('Reward'),
    ('Other');


SELECT COUNT(*) FROM transactions 
WHERE sender_id = 3 OR receiver_id = 3;

SELECT * FROM "transactions" 
WHERE "transactions"."deletd_at" IS NULL 
ORDER BY amount ASC 
LIMIT 25;

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

SELECT * FROM transactions;
SELECT * FROM games;

UPDATE games 
SET win_box_id = 1;

SELECT * FROM "transactions" WHERE (sender_id = 2 OR receiver_id = 2) AND "transactions"."deletd_at" IS NULL ORDER BY "amount" LIMIT 10
