-- =========================
-- TABELA: usuario
-- =========================
CREATE TABLE usuario (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(150) NOT NULL,
    email VARCHAR(150) UNIQUE NOT NULL,
    senha VARCHAR(255) NOT NULL,
    createdat TIMESTAMP DEFAULT NOW()
);

-- =========================
-- USUÁRIO DE TESTE
-- senha: 123456 (bcrypt)
-- =========================
INSERT INTO usuario (nome, email, senha)
VALUES (
    'Usuário Teste',
    'teste@teste.com',
    '$2a$10$uOraViEOfmAh2JBt.WqSCeAhzE4BdIviIX79wlpcwH5Jgyj2J2UZC'
);

-- =========================
-- ENUM: status da tarefa
-- =========================
CREATE TYPE status_tarefa AS ENUM ('A', 'F', 'C');
-- A = Aberta
-- F = Finalizada
-- C = Cancelada

-- =========================
-- TABELA: tarefa
-- =========================
CREATE TABLE tarefa (
    id SERIAL PRIMARY KEY,
    idusuario INT NOT NULL,
    titulo VARCHAR(200) NOT NULL,
    descricao VARCHAR(500),
    status status_tarefa NOT NULL DEFAULT 'A',
    dataagendada TIMESTAMP NULL,
    dataconclusao TIMESTAMP NULL,
    createdat TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_tarefas_usuario
        FOREIGN KEY (idusuario)
        REFERENCES usuario(id)
        ON DELETE CASCADE
);

-- =========================
-- INDEXES (performance)
-- =========================
CREATE INDEX idx_usuario_email ON usuario(email);
CREATE INDEX idx_tarefa_idusuario ON tarefa(idusuario);
