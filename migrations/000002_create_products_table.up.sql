-- Create products table
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    current_price INTEGER,
    url VARCHAR(250),
    user_id INTEGER
);

-- Add foreign key
ALTER TABLE products 
ADD CONSTRAINT fk_user_id
FOREIGN KEY (user_id)
REFERENCES users (id);

