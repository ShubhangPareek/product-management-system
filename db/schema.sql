CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    product_name VARCHAR(255) NOT NULL,
    product_description TEXT,
    product_images TEXT[],
    compressed_product_images TEXT[],
    product_price NUMERIC(10, 2) NOT NULL
);