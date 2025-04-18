-- Create transaction_types table
CREATE TABLE IF NOT EXISTS transaction_types (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create users table
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create transaction_categories table
CREATE TABLE IF NOT EXISTS transaction_categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create transactions table
CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    amount DECIMAL(10,2) NOT NULL,
    type INTEGER NOT NULL,
    category INTEGER NOT NULL,
    detail VARCHAR(500) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (type) REFERENCES transaction_types(id),
    FOREIGN KEY (category) REFERENCES transaction_categories(id)
);

-- Insert default transaction types
INSERT INTO transaction_types (name) VALUES 
('debit'),
('credit');

-- Insert default transaction categories
INSERT INTO transaction_categories (name) VALUES 
('food & drink'),
('bills'),
('salary'),
('rent'),
('groceries'),
('transportation'),
('travel'),
('shopping'),
('entertainment'),
('education'),
('family'),
('health'),
('other');