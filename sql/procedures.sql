CREATE OR REPLACE FUNCTION criartransacao(
  IN idcliente integer,
  IN valor integer,
  IN descricao varchar(10)
) RETURNS RECORD AS $$
DECLARE
  clienteencontrado clientes%rowtype;
  ret RECORD;
BEGIN
  SELECT * FROM clientes
  INTO clienteencontrado
  WHERE id = idcliente;

  IF not found THEN
    SELECT -1 INTO ret;
    RETURN ret;
  END IF;

  UPDATE clientes
    SET saldo = saldo + valor
    WHERE id = idcliente AND (valor > 0 OR saldo + valor >= limite)
    RETURNING saldo, limite
    INTO ret;
  IF ret.limite is NULL THEN
    SELECT -2 INTO ret;
    RETURN ret;
  END IF;
  INSERT INTO transacaos (valor, descricao, idcliente)
    VALUES (valor, descricao, cliente_id);
  RETURN ret;
END;$$ LANGUAGE plpgsql;