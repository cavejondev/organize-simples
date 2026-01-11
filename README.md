# organize-simples

App para controle de tarefas e finanças.

**Versão:** 1.0

---

## Pré-requisitos

- [Go 1.24+](https://go.dev/dl/)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Git](https://git-scm.com/)

---

## Instalação

1. Clone o repositório:

```bash
git clone https://github.com/cavejondev/organize-simples
cd organize-simples
```

2. Crie um banco de dados PostgreSQL e rode os comandos abaixo no banco, você pode mudar o email no insert, a senha ja vem padrao como "123456":

```sql
-- Tabela de usuários
CREATE TABLE usuario (
    id SERIAL PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    senha TEXT NOT NULL,
    createdat TIMESTAMP DEFAULT NOW()
);

-- Inserir usuário de teste
INSERT INTO usuario(email, senha)
VALUES ('teste@teste.com', '$2a$10$uOraViEOfmAh2JBt.WqSCeAhzE4BdIviIX79wlpcwH5Jgyj2J2UZC');

-- Enum para status das tarefas
CREATE TYPE status_tarefa AS ENUM ('A', 'F', 'C');

-- Tabela de tarefas
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
```

3. Configure o arquivo `.env` na raiz do projeto, pode ser o nome que quiser do banco, tambem a secret e expire do token, a porta tambem pode ser modificada:

```env
PORT=8080

DATABASE_URL=postgres://usuario:senha@localhost:5432/meubanco?sslmode=disable

JWT_SECRET=secret
JWT_EXPIRE_HOURS=24
```

4. Instale as dependências Go:

```bash
go mod tidy
```

---

## Execução

Execute o projeto:

```bash
go run ./cmd/main.go
```

O servidor iniciará na porta configurada no `.env` (padrão 8080).

---

## Rotas principais

### Autenticação

- `POST /login`

  - Body JSON:

  ```json
  {
    "email": "teste@teste.com",
    "senha": "123456"
  }
  ```

  - Retorna JWT para autenticação.

### Tarefas (protegidas, requerem token)

- `POST /tarefa` - Criar tarefa
- `GET /tarefa` - Listar tarefas do usuário
- `PUT /tarefa/{id}` - Atualizar tarefa

> ⚠️ Para todas as rotas de tarefas, enviar header `Authorization: Bearer <TOKEN>`.

---

## Estrutura do projeto

```
.
├── internal
│   ├── db           # Conexão com Postgres
│   ├── domain       # Models e interfaces
│   ├── handlers     # Handlers HTTP
│   ├── repositories # Implementação dos repositórios
│   ├── services     # Lógica de negócio
│   └── utils        # Funções utilitárias
├── go.mod
├── go.sum
├── main.go
└── .env
```

---

## Observações

- Projeto segue padrão de DI (Dependency Injection) para services e repositórios.
- JWT usado para autenticação.
- Status de tarefa:

  - `A` - Aberto
  - `F` - Finalizado
  - `C` - Cancelado

---

## Licença

MIT
