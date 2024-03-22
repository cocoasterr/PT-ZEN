-- Membuat tabel users
CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(255) PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    firstname VARCHAR,
    lastname VARCHAR,
    phone_number VARCHAR,
    password VARCHAR(255) NOT NULL,
    deleted_at TIMESTAMP
);

-- Membuat tabel product
CREATE TABLE IF NOT EXISTS product (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255),
    harga INT,
    stock INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(255),
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by VARCHAR(255),
    deleted_at TIMESTAMP
);

-- Membuat tabel pesanan
CREATE TABLE IF NOT EXISTS pesanan (
    id VARCHAR(255) PRIMARY KEY,
    product_id VARCHAR(255),
    user_id VARCHAR(255),
    qty INT,
    status BOOLEAN,
    deleted_at TIMESTAMP,
    CONSTRAINT pesanan_product_id_fkey FOREIGN KEY (product_id) REFERENCES product(id),
    CONSTRAINT pesanan_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id)
);