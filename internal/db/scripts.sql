CREATE TABLE usuario (
    id SERIAL PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    senha TEXT NOT NULL,
    createdat TIMESTAMP DEFAULT NOW()
);

INSERT INTO usuario(email, senha) 
VALUES ('teste@teste.com', '$2a$10$uOraViEOfmAh2JBt.WqSCeAhzE4BdIviIX79wlpcwH5Jgyj2J2UZC');

CREATE TYPE status_tarefa AS ENUM ('A', 'F', 'C');

CREATE TABLE tarefa (
    id SERIAL PRIMARY KEY,

    idusuario INT NOT NULL,

    titulo TEXT NOT NULL,
    descricao TEXT,

    status status_tarefa NOT NULL DEFAULT 'A',

    dataagendada TIMESTAMP NULL,
    dataconclusao TIMESTAMP NULL,

    createdat TIMESTAMP NOT NULL DEFAULT now(),

    CONSTRAINT fk_tarefas_usuario
        FOREIGN KEY (idusuario) REFERENCES usuario(id)
);
