DROP TABLE IF EXISTS razzzila.currencies;
DROP TABLE IF EXISTS razzzila.currencies_currencies;
DROP TABLE IF EXISTS razzzila.currencies_exchange_rates;
DROP TABLE IF EXISTS razzzila.currencies_exchange_rates_history;


CREATE TABLE IF NOT EXISTS razzzila.currencies_currencies
(
    id INT,
    code VARCHAR(255),
    name VARCHAR(255),
    symbol VARCHAR(10),
    decimal_number INT,
    active BOOLEAN,
    main_area_id INT,

    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS razzzila.currencies_exchange_rates
(
    id INT,
    currency_id INT,
    target_currency_id INT,
    exchange_rate FLOAT,
    rate_source_id INT,
    created_at DateTime,
    updated_at DateTime,

    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS razzzila.currencies_exchange_rates_history
(
    id INT,
    currency_id INT,
    target_currency_id INT,
    exchange_rate FLOAT,
    rate_source_id INT,
    update_date DateTime,

    PRIMARY KEY (id)
);


DROP TABLE IF EXISTS razzzila.testing;
CREATE TABLE IF NOT EXISTS razzzila.testing
(
    id INT,
    code VARCHAR(255),
    active BOOLEAN,
    main_area_id INT,

    PRIMARY KEY (id)
);