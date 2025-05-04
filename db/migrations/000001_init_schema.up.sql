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
('foodanddrink'),
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
(1, 25000, 1, 1, 'Lunch at Warung Makan', '2025-01-02'),
(1, 40000, 1, 2, 'Electricity bill', '2025-01-03'),
(1, 10000000, 2, 3, 'January Salary', '2025-01-05'),
(1, 200000, 1, 5, 'Groceries at Alfamart', '2025-01-08'),
(1, 5000000, 1, 4, 'Monthly rent', '2025-02-02'),
(1, 200000, 1, 5, 'Groceries at Superindo', '2025-02-04'),
(1, 15000, 1, 6, 'Train top-up', '2025-02-06'),
(1, 10000000, 2, 3, 'February Salary', '2025-02-10'),
(1, 300000, 1, 7, 'Train to Bandung', '2025-03-01'),
(1, 150000, 1, 9, 'Cinema ticket', '2025-03-05'),
(1, 75000, 1, 1, 'Dinner at Padang restaurant', '2025-03-10'),
(1, 6000000, 2, 3, 'Bonus payout', '2025-03-15'),
(1, 2000000, 1, 10, 'Tuition fee', '2025-04-01'),
(1, 1100000, 1, 11, 'Family monthly', '2025-04-05'),
(1, 150000, 1, 12, 'Pharmacy expense', '2025-04-06'),
(1, 10000000, 2, 3, 'April Salary', '2025-04-10'),
(1, 150000, 1, 1, 'Brunch with friends', '2025-05-01'),
(1, 5000000, 2, 3, 'Freelance project payment', '2025-05-02'),
(1, 350000, 1, 8, 'Bought Bluetooth headphones', '2025-05-04');