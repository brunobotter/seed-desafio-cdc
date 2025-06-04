-- Tabela Author
CREATE TABLE author (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_author_name (name)
);

-- Tabela Category
CREATE TABLE category (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_category_name (name)
);

-- Tabela Book
CREATE TABLE book (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    resume TEXT NOT NULL,
    summary TEXT NOT NULL,
    price DECIMAL(10,2) NOT NULL CHECK (price > 0),
    page INT NOT NULL CHECK (page > 0),
    isbn VARCHAR(20) NOT NULL UNIQUE,
    publish_date DATE NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    author_id INT NOT NULL,
    category_id INT NOT NULL,
    FOREIGN KEY (author_id) REFERENCES author(id),
    FOREIGN KEY (category_id) REFERENCES category(id),
    INDEX idx_book_title (title),
    INDEX idx_book_author (author_id),
    INDEX idx_book_category (category_id)
);

-- Tabela Country
CREATE TABLE country (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_country_name (name)
);

-- Tabela State
CREATE TABLE state (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    country_id INT NOT NULL,
    FOREIGN KEY (country_id) REFERENCES country(id),
    INDEX idx_state_name (name),
    INDEX idx_state_country (country_id)
);
