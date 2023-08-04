CREATE TABLE users(
    id INT AUTO_INCREMENT,
    email VARCHAR(255) NOT NULL UNIQUE,
    encrypted_password VARCHAR(255) NOT NULL,
    PRIMARY KEY(id)
);