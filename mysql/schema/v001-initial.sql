DROP TABLE IF EXISTS razzzila.currencies_currencies;
DROP TABLE IF EXISTS razzzila.currencies_exchange_rates;
DROP TABLE IF EXISTS razzzila.currencies_exchange_rates_history;

CREATE TABLE IF NOT EXISTS razzzila.currencies_currencies
(
    id INT,
    code VARCHAR(3),
    name VARCHAR(255),
    symbol VARCHAR(10),
    decimal_number SMALLINT,
    active BOOLEAN,
    main_area_id INT,

    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS razzzila.currencies_exchange_rates
(
    id INT,
    currency_id SMALLINT,
    target_currency_id SMALLINT,
    exchange_rate TINYINT,
    rate_source_id SMALLINT,
    created_at DateTime,
    updated_at DateTime,

    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS razzzila.currencies_exchange_rates_history
(
    id INT,
    currency_id SMALLINT,
    target_currency_id SMALLINT,
    exchange_rate TINYINT,
    rate_source_id SMALLINT,
    update_date DateTime,

    PRIMARY KEY (id)
);