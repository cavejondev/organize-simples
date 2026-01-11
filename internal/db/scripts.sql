CREATE TABLE usuario (
    id SERIAL PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    senha TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

INSERT INTO usuario(email, senha) 
VALUES ('teste@teste.com', '$2a$10$uOraViEOfmAh2JBt.WqSCeAhzE4BdIviIX79wlpcwH5Jgyj2J2UZC');