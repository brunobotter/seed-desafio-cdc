CREATE TABLE payments (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    customer_id BIGINT NOT NULL,
    total DECIMAL(10,2) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE payment_items (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    payment_id BIGINT NOT NULL,
    book_id BIGINT NOT NULL,
    amount INT NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    FOREIGN KEY (payment_id) REFERENCES payments(id) ON DELETE CASCADE
);