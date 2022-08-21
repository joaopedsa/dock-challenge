CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY NOT NULL,
    name TEXT NOT NULL,
    cpf TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS bank_accounts (
    id SERIAL PRIMARY KEY NOT NULL,
    user_id INTEGER NOT NULL,
    balance NUMERIC(38, 18) NOT NULL DEFAULT 0,
    number TEXT NOT NULL,
    agency TEXT NOT NULL,
    is_active BOOL NOT NULL DEFAULT true,
    is_blocked BOOL NOT NULL DEFAULT false,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE RESTRICT,
    UNIQUE(number, agency)
);

CREATE TYPE type_statement AS ENUM (
    'WITHDRAW', 'DEPOSIT'
);

CREATE TABLE IF NOT EXISTS bank_statements (
    id SERIAL PRIMARY KEY NOT NULL,
    bank_accounts_id INTEGER NOT NULL,
    value NUMERIC(38, 18) NOT NULL DEFAULT 0,
    type type_statement NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (bank_accounts_id) REFERENCES bank_accounts (id) ON DELETE RESTRICT
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON bank_accounts
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON bank_statements
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

ALTER TABLE bank_accounts ADD CONSTRAINT validate_balance CHECK (balance >= 0);
