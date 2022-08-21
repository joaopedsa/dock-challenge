INSERT INTO users (id, name, cpf) VALUES (1, 'João Pedro Santana', '09662672907');
INSERT INTO bank_accounts (id, user_id, balance, number, agency) VALUES (1, 1, 1000, '12345678', '0001');
INSERT INTO bank_statements (bank_accounts_id, value, type) VALUES (1, 1500, 'DEPOSIT');
INSERT INTO bank_statements (bank_accounts_id, value, type) VALUES (1, 500, 'WITHDRAW');

INSERT INTO users (id, name, cpf) VALUES (2, 'Gustavo Ricardo Santana', '749.134.769-04');
INSERT INTO bank_accounts (id, user_id, balance, number, agency) VALUES (2, 2, 500, '12345678', '0002');
INSERT INTO bank_statements (bank_accounts_id, value, type) VALUES (2, 1000, 'DEPOSIT');
INSERT INTO bank_statements (bank_accounts_id, value, type) VALUES (2, 500, 'WITHDRAW');
