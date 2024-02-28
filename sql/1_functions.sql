CREATE OR REPLACE FUNCTION criatransacao(
  IN idcliente integer,
  IN tipo char(1),
  IN valor integer,
  IN descricao varchar(10)
) RETURNS RECORD AS $$
DECLARE
  clienteencontrado cliente%rowtype;
  ret RECORD;
BEGIN
  SELECT * FROM cliente
  INTO clienteencontrado
  WHERE id = idcliente;

  IF not found THEN
    SELECT -1 INTO ret;
    RETURN ret;
  END IF;

  CASE 
    WHEN tipo = 'd' THEN
      UPDATE cliente
        SET saldo = saldo - valor
        WHERE id = idcliente AND (valor > 0 OR ABS(saldo - valor) >= limite)
        RETURNING saldo, limite
        INTO ret;
      IF ret.limite is NULL THEN
        SELECT -2 INTO ret;
        RETURN ret;
      END IF;
    WHEN tipo = 'c' THEN
      UPDATE cliente
        SET saldo = saldo + valor
        WHERE id = idcliente
        RETURNING saldo, limite
        INTO ret;
    ELSE
      SELECT -1 INTO ret;
      RETURN ret;
  END CASE;
  INSERT INTO transacao (tipo, valor, descricao, cliente_id)
    VALUES (tipo, valor, descricao, idcliente);
  RETURN ret;
END;$$ LANGUAGE plpgsql;