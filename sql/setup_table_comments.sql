CREATE TABLE comments2 (
    id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    content TEXT,
    ceate_date DATETIME NOT NULL DEFAULT AUTO_INCREMENT CURRENT_TIMESTAMP,
    task_id INT
);