-- +goose Up
CREATE TABLE cadastros_imobiliarios (
  objectid serial PRIMARY KEY,
  codigo_cadastro integer NOT NULL UNIQUE,
  tipo_cadastro integer NOT NULL,
  data_cadastro date NOT NULL,
  codigo_terreno integer,
  codigo_cadastro_global varchar,
  situacao_cadastral integer NOT NULL,
  inscricao_imobiliaria varchar NOT NULL,
  inscricao_anterior varchar,
  observacao varchar,
  unico_cartorio int,
  matricula varchar,
  nro_livro varchar,
  nro_folha varchar,
  area_terreno float NOT NULL,
  area_terreno_escriturada float NOT NULL,
  profundidade float NOT NULL,
  area_construida float NOT NULL,
  area_construida_averbada float NOT NULL,
  area_total_construida float NOT NULL,
  area_comum float NOT NULL,
  afastamento_frontal float NOT NULL,
  numero_pavimentos integer NOT NULL,
  nome_propriedade varchar,
  numero_incra varchar,
  numero_receita_federal varchar,
  area_terreno_rural float NOT NULL,
  area_terreno_escriturada_rural float NOT NULL,
  area_construida_rural float NOT NULL,
  area_construida_averbada_rural float NOT NULL,
  tipo_medida int,
  economicos varchar,
  data_hora_ultima_alteracao timestamp NOT NULL,
  globalid varchar NOT NULL UNIQUE,
  created_date timestamp NOT NULL,
  last_edited_date timestamp NOT NULL
);

CREATE TABLE caracteristicas (
  objectid serial PRIMARY KEY,
  codigo_cadastro integer NOT NULL,
  codigo_bloco integer NOT NULL,
  codigo_item integer NOT NULL,
  sequencia_item integer NOT NULL,
  valor varchar,
  valor_lista integer,
  globalid varchar NOT NULL UNIQUE,
  parentglobalid varchar NOT NULL,
  created_date timestamp NOT NULL,
  last_edited_date timestamp NOT NULL
);

CREATE TABLE subreceitas (
  objectid serial PRIMARY KEY,
  codigo_cadastro integer NOT NULL,
  codigo_subreceita integer NOT NULL,
  data_inicio_vigencia date NOT NULL,
  data_fim_vigencia date,
  situacao integer NOT NULL,
  globalid varchar NOT NULL UNIQUE,
  parentglobalid varchar NOT NULL,
  created_date timestamp NOT NULL,
  last_edited_date timestamp NOT NULL
);

CREATE TABLE enderecos (
  objectid serial PRIMARY KEY,
  codigo_cadastro integer NOT NULL,
  tipo_endereco integer NOT NULL,
  descricao_cidade varchar NOT NULL,
  descricao_bairro varchar NOT NULL,
  descricao_logradouro varchar NOT NULL,
  numero varchar,
  inf_complementar varchar,
  complemento varchar,
  garagem varchar,
  sala varchar,
  loja varchar,
  quadra varchar,
  lote varchar,
  bloco varchar,
  coordenada_latitude float,
  coordenada_longitude float,
  coordenada_panorama float,
  codigo_cidade integer NOT NULL,
  codigo_bairro integer NOT NULL,
  codigo_logradouro integer NOT NULL,
  cep integer NOT NULL,
  codigo_loteamento integer,
  codigo_edificio integer,
  nro_apto integer,
  globalid varchar NOT NULL UNIQUE,
  parentglobalid varchar NOT NULL,
  created_date timestamp NOT NULL,
  last_edited_date timestamp NOT NULL
);

CREATE TABLE proprietarios (
  objectid serial PRIMARY KEY,
  codigo_cadastro integer NOT NULL,
  tipo_proprietario integer NOT NULL,
  situacao integer NOT NULL,
  percentual float NOT NULL,
  data_inicio_vigencia timestamp NOT NULL,
  codigo integer NOT NULL,
  globalid varchar NOT NULL UNIQUE,
  parentglobalid varchar NOT NULL,
  created_date timestamp NOT NULL,
  last_edited_date timestamp NOT NULL
);

CREATE TABLE subtipoproprietarios (
  objectid SERIAL PRIMARY KEY,
  codigo_cadastro integer NOT NULL,
  codigo_proprietario integer NOT NULL,
  codigo integer NOT NULL,
  tipo integer NOT NULL,
  descricao varchar NOT NULL,
  situacao varchar NOT NULL,
  globalid varchar NOT NULL UNIQUE,
  parentglobalid varchar NOT NULL,
  created_date timestamp NOT NULL,
  last_edited_date timestamp NOT NULL
);

CREATE TABLE testadas (
  objectid serial PRIMARY KEY,
  codigo_cadastro integer NOT NULL,
  numero_testada integer NOT NULL,
  metragem float NOT NULL,
  codigo_secao integer NOT NULL,
  id_secao varchar,
  globalid varchar NOT NULL UNIQUE,
  parentglobalid varchar NOT NULL,
  created_date timestamp NOT NULL,
  last_edited_date timestamp NOT NULL
);

CREATE TABLE zoneamentos (
  objectid serial PRIMARY KEY,
  codigo_cadastro integer NOT NULL,
  descricao varchar NOT NULL,
  observacao varchar,
  codigo_zoneamento integer NOT NULL,
  principal integer NOT NULL,
  globalid varchar NOT NULL UNIQUE,
  parentglobalid varchar NOT NULL,
  created_date timestamp NOT NULL,
  last_edited_date timestamp NOT NULL
);

-- +goose Down
DROP TABLE cadastros_imobiliarios;
DROP TABLE caracteristicas;
DROP TABLE enderecos;
DROP TABLE proprietarios;
DROP TABLE testadas;
DROP TABLE zoneamentos;
DROP TABLE subreceitas;
DROP TABLE subtipoproprietarios;
