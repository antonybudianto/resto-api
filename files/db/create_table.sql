CREATE DATABASE restohub;
USE restohub;

CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(320) NOT NULL
);

CREATE TABLE cuisines (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

CREATE TABLE countries (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

CREATE TABLE restaurants (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    slug VARCHAR(50) NOT NULL,
    cuisine_id INT NOT NULL,
    country_id INT NOT NULL,
    lat DECIMAL(10, 8) NOT NULL,
    lng DECIMAL(11, 8) NOT NULL,
    address TEXT,
    rating DECIMAL(2, 1),

    CONSTRAINT FK_RESTO_CUISINE FOREIGN KEY (cuisine_id) REFERENCES cuisines(id),
    CONSTRAINT FK_RESTO_COUNTRY FOREIGN KEY (country_id) REFERENCES countries(id)
);

CREATE TABLE books (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    restaurant_id INT NOT NULL,
    book_datetime TIMESTAMP NOT NULL,
    total_people TINYINT NOT NULL,

    CONSTRAINT FK_BOOK_USER FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT FK_BOOK_RESTO FOREIGN KEY (restaurant_id) REFERENCES restaurants(id)
);
