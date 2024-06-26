CREATE TABLE IF NOT EXISTS testdb.currencies 
(
    code VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255),
    symbol_native VARCHAR(10),
    decimal_digits INT,
    active BOOLEAN,
    main_area_id INT,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME
);

INSERT INTO testdb.currencies (code, name, symbol_native, decimal_digits, active, main_area_id) VALUES ('USD', 'US Dollar', '$', 2, 0, 0, NOW(), NOW(), NULL), ('EUR', 'Euro', '€', 2, 0, 0, NOW(), NOW(), NULL);

