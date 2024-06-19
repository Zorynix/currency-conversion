DROP TABLE IF EXISTS razzzila.currencies;
DROP TABLE IF EXISTS razzzila.exchange_rates;
DROP TABLE IF EXISTS razzzila.exchange_rate_histories;


CREATE TABLE IF NOT EXISTS razzzila.currencies
(
    code VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255),
    symbol_native VARCHAR(10),
    decimal_digits INT,
    active BOOLEAN,
    main_area_id INT
);

CREATE TABLE IF NOT EXISTS razzzila.exchange_rates
(
    code VARCHAR(255) PRIMARY KEY,
    currency_id INT,
    target_currency_id INT,
    exchange_rate FLOAT,
    rate_source_id INT,
    created_at DateTime,
    updated_at DateTime,
);

CREATE TABLE IF NOT EXISTS razzzila.exchange_rate_histories
(
    code VARCHAR(255) PRIMARY KEY,
    currency_id INT,
    target_currency_id INT,
    exchange_rate FLOAT,
    rate_source_id INT,
    created_at DateTime,
    updated_at DateTime,
);