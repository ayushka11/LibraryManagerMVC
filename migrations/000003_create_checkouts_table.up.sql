CREATE TABLE IF NOT EXISTS checkouts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    book_id INT,
    checkout_date DATE,
    due_date DATE,
    return_date DATE,
    fine DECIMAL(5,2) DEFAULT 0,
    status ENUM('pending', 'approved', 'rejected') DEFAULT 'pending',
    type ENUM('checkout', 'checkin'),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (book_id) REFERENCES books(id)
);
