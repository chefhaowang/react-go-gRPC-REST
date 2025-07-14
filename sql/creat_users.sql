CREATE DATABASE IF NOT EXISTS users;
USE users;


CREATE TABLE users (
    id CHAR(36) PRIMARY KEY,                -- UUID
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    role ENUM('admin', 'customer') NOT NULL DEFAULT 'customer',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


-- Admin user
INSERT INTO users (id, email, password_hash, role) VALUES (
  UUID(), 
  'admin@example.com', 
  '$2a$10$CE0x1bGEY0wHq9eSg9Erw.kZmCOrcEYZx8wNZYtvQzoih.XbPCFxW', 
  'admin'
);

-- Customer users
INSERT INTO users (id, email, password_hash, role) VALUES
(
  UUID(), 
  'alice@example.com', 
  '$2a$10$W5vuWOP7U5V9JmyKzq7xleT6Uoz3aEvS0uQW3wFQHfcdOxOMExiIS', 
  'customer'
),
(
  UUID(), 
  'bob@example.com', 
  '$2a$10$W5vuWOP7U5V9JmyKzq7xleT6Uoz3aEvS0uQW3wFQHfcdOxOMExiIS', 
  'customer'
);