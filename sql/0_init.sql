
DROP TABLE IF EXISTS transacaos;
DROP TABLE IF EXISTS clientes;


CREATE TABLE transacaos (
    id SERIAL PRIMARY KEY,
    cliente_id BIGINT NOT NULL,
    valor BIGINT NOT NULL,
    tipo CHAR(1) NOT NULL,
    descricao VARCHAR(128) NOT NULL,
    criado_em TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE clientes (
    id SERIAL PRIMARY KEY,
    limite BIGINT NOT NULL,
    saldo BIGINT NOT NULL DEFAULT 0
);

DO $$
BEGIN
  INSERT INTO clientes (limite)
  VALUES
    (100000),
    (80000),
    (1000000),
    (10000000),
    (500000);
END; $$


