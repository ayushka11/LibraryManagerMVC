CREATE TABLE IF NOT EXISTS books (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    available INT DEFAULT 1,
    quantity INT DEFAULT 1,
    CONSTRAINT non_negative_constraint CHECK (available >= 0),
    CONSTRAINT available_less_than_quantity CHECK (available <= quantity)
);

