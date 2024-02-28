DROP TABLE IF EXISTS cliente;
DROP TABLE IF EXISTS transacao;


CREATE TABLE cliente (
    id SERIAL PRIMARY KEY,
    limite BIGINT NOT NULL,
    saldo INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE transacao (
    id SERIAL PRIMARY KEY,
    cliente_id INTEGER NOT NULL,
    valor BIGINT NOT NULL,
    tipo CHAR(1) NOT NULL,
    descricao VARCHAR(128) NOT NULL,
    criado_em TIMESTAMP NOT NULL DEFAULT now(),
    CONSTRAINT fk_cliente_transacao FOREIGN KEY(cliente_id) REFERENCES cliente(id)
);



DO $$
BEGIN
  INSERT INTO cliente (limite)
  VALUES
    (100000),
    (80000),
    (1000000),
    (10000000),
    (500000);
END; $$


