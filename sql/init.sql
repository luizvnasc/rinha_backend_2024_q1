CREATE TABLE transacaos (
    id SERIAL PRIMARY KEY,
    cliente_id BIGINT NOT NULL,
    valor BIGINT NOT NULL,
    tipo CHAR(1) NOT NULL,
    descricao VARCHAR(128) NOT NULL,
    created_at TIMESTAMP NOT NULL
)