CREATE TABLE customer (
    id INT AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    document VARCHAR(20) NOT NULL,
    address TEXT NOT NULL,
    complement TEXT NOT NULL,
    city VARCHAR(100) NOT NULL,
    country_id INT NOT NULL,
    state_id INT DEFAULT NULL,
    phone VARCHAR(20) NOT NULL,
    cep VARCHAR(20) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (country_id) REFERENCES country(id),
    FOREIGN KEY (state_id) REFERENCES state(id),

    UNIQUE (email),
    UNIQUE (document),

    INDEX idx_customer_email (email),
    INDEX idx_customer_document (document),
    INDEX idx_customer_country (country_id),
    INDEX idx_customer_state (state_id)
);