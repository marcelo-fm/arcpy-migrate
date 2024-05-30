-- +goose Up
CREATE VIEW vw_dados_terreno AS
SELECT
  objectid,
  codigo_cadastro,
  tipo_cadastro,
  codigo_terreno,
  codigo_cadastro_global,
  situacao_cadastral,
  inscricao_imobiliaria,
  inscricao_anterior,
  observacao,
  unico_cartorio,
  nro_livro,
  nro_folha,
  area_terreno,
  area_terreno_escriturada,
  profundidade,
  globalid,
  created_date,
  last_edited_date
FROM
  cadastros_imobiliarios
WHERE
  tipo_cadastro = 1;

CREATE VIEW vw_dados_unidade AS
SELECT
  objectid,
  codigo_cadastro,
  tipo_cadastro,
  codigo_terreno,
  codigo_cadastro_global,
  situacao_cadastral,
  inscricao_imobiliaria,
  inscricao_anterior,
  observacao,
  unico_cartorio,
  nro_livro,
  nro_folha,
  area_construida,
  area_construida_averbada,
  area_total_construida,
  area_comum,
  afastamento_frontal,
  numero_pavimentos,
  globalid,
  created_date,
  last_edited_date
FROM
  cadastros_imobiliarios
WHERE
  tipo_cadastro = 2;

CREATE VIEW vw_dados_rural AS
SELECT
  objectid,
  codigo_cadastro,
  tipo_cadastro,
  codigo_terreno,
  codigo_cadastro_global,
  situacao_cadastral,
  inscricao_imobiliaria,
  inscricao_anterior,
  observacao,
  unico_cartorio,
  nro_livro,
  nro_folha,
  nome_propriedade,
  numero_incra,
  numero_receita_federal,
  area_terreno_rural,
  area_terreno_escriturada_rural,
  area_construida_rural,
  area_construida_averbada_rural,
  tipo_medida,
  economicos,
  globalid,
  created_date,
  last_edited_date
FROM
  cadastros_imobiliarios
WHERE
  tipo_cadastro = 3;

-- +goose Down
DROP VIEW vw_dados_terreno;
DROP VIEW vw_dados_unidade;
DROP VIEW vw_dados_rural;
