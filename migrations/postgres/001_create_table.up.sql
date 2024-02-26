CREATE TABLE IF NOT EXISTS author (
    id UUID PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    login VARCHAR(50) NOT NULL,
    password VARCHAR(128) NOT NULL,
    phone VARCHAR(13) NOT NULL,
    gender VARCHAR(15) CHECK (gender IN('male', 'female')),
    birth_date DATE NOT NULL,
    age INT,
    address VARCHAR(120) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at INTEGER DEFAULT 0
  );
CREATE TABLE IF NOT EXISTS books (
    id UUID PRIMARY KEY NOT NULL,
    name VARCHAR(50) NOT NULL,
    author_name VARCHAR(50) NOT NULL,
    page_number INT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at INTEGER DEFAULT 0
);