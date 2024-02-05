DROP TABLE IF EXISTS razzzila.currencies;
DROP TABLE IF EXISTS razzzila.currencies_exchange_rates;
DROP TABLE IF EXISTS razzzila.currency_exchange_rate_histories;
DROP TABLE IF EXISTS razzzila.currencies_exchange_rate_histories;


CREATE TABLE IF NOT EXISTS razzzila.currencies
(
    id INT AUTO_INCREMENT PRIMARY KEY,
    code VARCHAR(255),
    name VARCHAR(255),
    symbol VARCHAR(10),
    decimal_number INT,
    active BOOLEAN,
    main_area_id INT
);

CREATE TABLE IF NOT EXISTS razzzila.currencies_exchange_rates
(
    id INT AUTO_INCREMENT PRIMARY KEY,
    currency_id INT,
    target_currency_id INT,
    exchange_rate FLOAT,
    rate_source_id INT,
    created_at DateTime,
    updated_at DateTime
);

CREATE TABLE IF NOT EXISTS razzzila.currencies_exchange_rate_histories
(
    id INT AUTO_INCREMENT PRIMARY KEY,
    currency_id INT,
    target_currency_id INT,
    exchange_rate FLOAT,
    rate_source_id INT,
    update_date DateTime
);