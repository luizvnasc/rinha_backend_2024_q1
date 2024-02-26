CREATE TABLE transaction {
    id BIGINT NOT NULL,
    cliente_id BIGINT NOT NULL,
    valor BIGINT NOT NULL,
    tipo CHAR(1) NOT NULL,
    descricao VARCHAR(128) NOT NULL,
    criada_em DATE NOT NULL
}