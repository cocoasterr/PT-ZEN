-- Membuat tabel users
CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(255) PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    firstname VARCHAR,
    lastname VARCHAR,
    phone_number VARCHAR,
    password VARCHAR(255) NOT NULL
);

-- Membuat tabel product
CREATE TABLE IF NOT EXISTS product (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255),
    harga INT,
    stock INT
);

-- Membuat tabel pesanan
CREATE TABLE IF NOT EXISTS pesanan (
    id SERIAL PRIMARY KEY,
    product_id VARCHAR(255),
    user_id VARCHAR(255),
    status BOOLEAN,
    FOREIGN KEY (product_id) REFERENCES product(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);