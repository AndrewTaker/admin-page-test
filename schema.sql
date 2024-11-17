CREATE TABLE IF NOT EXISTS admin_main (
    id INT AUTO_INCREMENT PRIMARY KEY,
    range_from INT NOT NULL,
    range_to INT NOT NULL,
    days_columns TEXT NOT NULL
);
