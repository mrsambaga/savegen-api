-- Create users table
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create transaction_types table
CREATE TABLE IF NOT EXISTS transaction_types (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
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
    date DATE NOT NULL,
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

INSERT INTO users (username, email) VALUES 
('Sam Wilson', 'sam@gmail.com'),
('Bucky Barnes', 'bucky@gmail.com'),
('Natasha Romanoff', 'natasha@gmail.com'),
('Bruce Banner', 'bruce@gmail.com'),
('Thor Odinson', 'thor@gmail.com');

INSERT INTO transactions (user_id, amount, type, category, detail, date) VALUES 
(1, 20000, 1, 1, 'Ketoprak', '2025-01-01'),
(1, 50000, 1, 2, 'Mobile internet data', '2025-01-02'),
(1, 10000000, 2, 3, '', '2025-01-03'),
(1, 5000000, 1, 4, 'Monthly kos', '2025-02-04'),
(1, 200000, 1, 5, 'Weekly groceries', '2025-02-05'),
(1, 15000, 1, 6, 'Topup Transjakarta', '2025-02-06'),
(1, 300000, 1, 7, 'Ticket to yogyakarta', '2025-03-07'),
(1, 500000, 2, 8, 'Adidas running shoes', '2025-04-08'),
(1, 150000, 1, 9, 'Watch movie in XXI', '2025-04-09'),
(1, 10000000, 2, 3, '', '2025-04-10'),
(1, 2000000, 1, 10, 'College Tuition', '2025-04-10'),
(1, 1100000, 1, 11, 'Monthly family expense', '2025-04-11'),
(1, 150000, 1, 12, 'Vitamin pills', '2025-04-12'),
(1, 7000, 1, 13, 'Parking fee', '2025-04-13');