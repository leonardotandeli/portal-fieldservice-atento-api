DROP DATABASE IF EXISTS INTRAFIELD1; 
CREATE DATABASE IF NOT EXISTS INTRAFIELD1;
USE INTRAFIELD1;

DROP TABLE IF EXISTS USUARIOS;
DROP TABLE IF EXISTS BDC_POSTS;
DROP TABLE IF EXISTS BDC_CATEGORIAS;
DROP TABLE IF EXISTS MAPA_OPERACIONAL;
DROP TABLE IF EXISTS SITES;
DROP TABLE IF EXISTS CLIENTES;
DROP TABLE IF EXISTS DOMINIOS;
DROP TABLE IF EXISTS DACS;

CREATE TABLE USUARIOS (
  IDUSUARIO INT(11) PRIMARY KEY NOT NULL AUTO_INCREMENT,
  NOME VARCHAR(100) NOT NULL,
  LOGIN_NT VARCHAR(50) NOT NULL,
  RE VARCHAR(50) UNIQUE KEY NOT NULL,
  CARGO ENUM('ADM', 'M4', 'M3', 'M2', 'M1', 'P4', 'P3', 'P2', 'P1', 'N2') NOT NULL,
  EMAIL VARCHAR(50) UNIQUE KEY NOT NULL,
  SENHA VARCHAR(100) NOT NULL,
  V_USUARIOS ENUM('S','N') NOT NULL,
  V_BDC_POSTS ENUM('S','N') NOT NULL,
  V_BDC_ADM ENUM('S','N') NOT NULL,
  V_IMDB ENUM('S','N') NOT NULL,
  V_GSA ENUM('S','N') NOT NULL,
  V_MAPA_OPERACIONAL ENUM('S','N') NOT NULL,
  V_MAPA_OPERACIONAL_ADM ENUM('S','N') NOT NULL,
  ID_SITE INT,
  DATA_CRIACAO TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP()
);

CREATE TABLE BDC_POSTS (
  IDPOST INT(11) PRIMARY KEY NOT NULL AUTO_INCREMENT,
  TITULO VARCHAR(100) NOT NULL,
  CONTEUDO LONGTEXT NOT NULL,
  ID_CATEGORIA INT,
  ID_USUARIO INT,
  ID_SITE INT,
  ID_CLIENTE INT, 
  DATA_CRIACAO TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP()
);

CREATE TABLE BDC_CATEGORIAS (
  IDCATEGORIA INT(11) PRIMARY KEY NOT NULL AUTO_INCREMENT,
  NOME VARCHAR(50) NOT NULL
);

CREATE TABLE MAPA_OPERACIONAL(
    IDMAPA INT AUTO_INCREMENT PRIMARY KEY,
    OPERACAO VARCHAR(100) NOT NULL,
    VLAN_DADOS VARCHAR(100) NOT NULL,
    VLAN_VOZ VARCHAR(100) NOT NULL,
    CONFIG_CONTRATUAL VARCHAR(100) NOT NULL,
    VERSAO_WINDOWS VARCHAR(100) NOT NULL,
    IMAGEM VARCHAR(100) NOT NULL,
    TEMPLATE LONGTEXT NOT NULL,
    GRUPO_IMDB VARCHAR(100) NOT NULL,
    GRAVADOR VARCHAR(200) NOT NULL,
    OBSERVACOES LONGTEXT NOT NULL,
    ID_SITE INT,
    ID_CLIENTE INT,
    ID_DOMINIO INT,
    ID_DAC INT,
    DATA_CRIACAO TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP()
);

CREATE TABLE SITES(
    IDSITE INT AUTO_INCREMENT PRIMARY KEY,
    NOME VARCHAR(100) NOT NULL
);
CREATE TABLE CLIENTES(
    IDCLIENTE INT AUTO_INCREMENT PRIMARY KEY,
    NOME VARCHAR(100) NOT NULL,
    LOGO_URL varchar(100) NOT NULL
);

CREATE TABLE DOMINIOS(
    IDDOMINIO INT AUTO_INCREMENT PRIMARY KEY,
    NOME VARCHAR(100) NOT NULL
);

CREATE TABLE DACS(
    IDDAC INT AUTO_INCREMENT PRIMARY KEY,
    NOME VARCHAR(100) NOT NULL
);

CREATE TABLE LOGS (
  IDUSUARIO INT(11) NOT NULL,
  LOGIN_NT VARCHAR(100) NOT NULL,
  NOME VARCHAR(100) NOT NULL,
  ACTION LONGTEXT NOT NULL,
  DATA TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP()
);

CREATE TABLE SESSIONS (
  ID INT(11) PRIMARY KEY NOT NULL AUTO_INCREMENT,
  ID_USUARIO INT(11) NOT NULL,
  TOKEN LONGTEXT NOT NULL,
  DATA TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP()

);

ALTER TABLE USUARIOS
ADD CONSTRAINT FK_USUARIOS_SITES
FOREIGN KEY(ID_SITE)
REFERENCES SITES(IDSITE);

ALTER TABLE BDC_POSTS
ADD CONSTRAINT FK_BDC_CATEGORIAS
FOREIGN KEY(ID_CATEGORIA)
REFERENCES BDC_CATEGORIAS(IDCATEGORIA);

ALTER TABLE BDC_POSTS
ADD CONSTRAINT FK_BDC_USUARIOS
FOREIGN KEY(ID_USUARIO)
REFERENCES USUARIOS(IDUSUARIO);

ALTER TABLE BDC_POSTS
ADD CONSTRAINT FK_BDC_SITES
FOREIGN KEY(ID_SITE)
REFERENCES SITES(IDSITE);

ALTER TABLE BDC_POSTS
ADD CONSTRAINT FK_BDC_CLIENTE
FOREIGN KEY(ID_CLIENTE)
REFERENCES CLIENTES(IDCLIENTE);

ALTER TABLE MAPA_OPERACIONAL
ADD CONSTRAINT FK_MAPA_SITE
FOREIGN KEY(ID_SITE)
REFERENCES SITES(IDSITE);

ALTER TABLE MAPA_OPERACIONAL
ADD CONSTRAINT FK_MAPA_CLIENTE
FOREIGN KEY(ID_CLIENTE)
REFERENCES CLIENTES(IDCLIENTE);

ALTER TABLE MAPA_OPERACIONAL
ADD CONSTRAINT FK_MAPA_DOMINIO
FOREIGN KEY(ID_DOMINIO)
REFERENCES DOMINIOS(IDDOMINIO);

ALTER TABLE MAPA_OPERACIONAL
ADD CONSTRAINT FK_MAPA_DAC
FOREIGN KEY(ID_DAC)
REFERENCES DACS(IDDAC);

ALTER TABLE SESSIONS
ADD CONSTRAINT FK_USER_SESSION
FOREIGN KEY(ID_USUARIO)
REFERENCES USUARIOS(IDUSUARIO);

INSERT INTO CLIENTES (IDCLIENTE, NOME, LOGO_URL) VALUES
	(1, 'INSTRUÇÕES GERAIS', '/assets/images/icons_clientes/geral.jpg'),
	(2, 'ALELO', '/assets/images/icons_clientes/alelo.jpg'),
	(3, 'APPLE', '/assets/images/icons_clientes/apple.jpg'),
	(4, 'ASURION', '/assets/images/icons_clientes/asurion.jpg'),
	(5, 'C6 BANK', '/assets/images/icons_clientes/c6.jpg'),
	(6, 'BANCO CARREFOUR', '/assets/images/icons_clientes/banco_carrefour.jpg'),
	(7, 'BANCO DO BRASIL', '/assets/images/icons_clientes/bb.jpg'),
	(8, 'BANCO PAN', '/assets/images/icons_clientes/banco_pan.jpg'),
	(9, 'BANCO SAFRA', '/assets/images/icons_clientes/safra.jpg'),
	(10, 'BANCO SOFISA', '/assets/images/icons_clientes/sofisa.jpg'),
	(11, 'BEMATECH', '/assets/images/icons_clientes/bematech.jpg'),
	(12, 'BENEFICIO FACIL', '/assets/images/icons_clientes/beneficiofacil.jpg'),
	(13, 'BMB', '/assets/images/icons_clientes/bmb.jpg'),
	(14, 'BMG', '/assets/images/icons_clientes/bmg.jpg'),
	(15, 'BOTICARIO', '/assets/images/icons_clientes/boticario.jpg'),
	(16, 'BRADESCO', '/assets/images/icons_clientes/bradesco.jpg'),
	(17, 'BURGER KING', '/assets/images/icons_clientes/bk.jpg'),
	(18, 'CAIXA SEGURADORA', '/assets/images/icons_clientes/caixa_seguradora.jpg'),
	(19, 'CATENO', '/assets/images/icons_clientes/cateno.jpg'),
	(20, 'CETELEM', '/assets/images/icons_clientes/cetelem.jpg'),
	(21, 'CIELO', '/assets/images/icons_clientes/cielo.jpg'),
	(22, 'CLARO', '/assets/images/icons_clientes/claro.jpg'),
	(23, 'CONECTCAR', '/assets/images/icons_clientes/connectcar.jpg'),
	(24, 'CONSULTING HOUSE', '/assets/images/icons_clientes/ch.jpg'),
	(25, 'CTF', '/assets/images/icons_clientes/ctf.jpg'),
	(26, 'DASA', '/assets/images/icons_clientes/dasa.jpg'),
	(27, 'DECOLAR', '/assets/images/icons_clientes/decolar.jpg'),
	(28, 'DEMAIS AREAS', '/assets/images/icons_clientes/demais_aereas.jpg'),
	(29, 'DISNEY', '/assets/images/icons_clientes/disney.jpg'),
	(30, 'EASYNVEST', '/assets/images/icons_clientes/easynvest.jpg'),
	(31, 'EDP', '/assets/images/icons_clientes/edp.jpg'),
	(32, 'EDITORA GLOBO', '/assets/images/icons_clientes/editora_globo.jpg'),
	(33, 'ENEL', '/assets/images/icons_clientes/enel.jpg'),
	(34, 'FACEBOOK', '/assets/images/icons_clientes/facebook.jpg'),
	(35, 'FIAT', '/assets/images/icons_clientes/fiat.jpg'),
	(36, 'FIRST DATA', '/assets/images/icons_clientes/firstdata.jpg'),
	(37, 'FORD BRASIL', '/assets/images/icons_clientes/ford.jpg'),
	(38, 'GOOGLE BRASIL', '/assets/images/icons_clientes/google.jpg'),
	(39, 'GPA', '/assets/images/icons_clientes/gpa.jpg'),
	(40, 'GRUPO ALIANCA', '/assets/images/icons_clientes/grupo_alianca.jpg'),
	(42, 'HONG', '/assets/images/icons_clientes/hong.jpg'),
	(43, 'HUAWEI', '/assets/images/icons_clientes/huawei.jpg'),
	(44, 'ICATU', '/assets/images/icons_clientes/icatu.jpg'),
	(45, 'IFOOD', '/assets/images/icons_clientes/ifood.jpg'),
	(46, 'INTERMEDICA', '/assets/images/icons_clientes/intermedica.jpg'),
	(47, 'INTERODONTO', '/assets/images/icons_clientes/interodonto.jpg'),
	(48, 'ITAU', '/assets/images/icons_clientes/itau.jpg'),
	(49, 'KROTON', '/assets/images/icons_clientes/kroton.jpg'),
	(50, 'LENOVO', '/assets/images/icons_clientes/lenovo.jpg'),
	(51, 'LIVELO', '/assets/images/icons_clientes/livelo.jpg'),
	(52, 'LOSANGO', '/assets/images/icons_clientes/losango.jpg'),
	(53, 'MARISA', '/assets/images/icons_clientes/marisa.jpg'),
	(54, 'MERCADO PAGO', '/assets/images/icons_clientes/mercadopago.jpg'),
	(55, 'MOTOROLA', '/assets/images/icons_clientes/motorola.jpg'),
	(56, 'NESTLE', '/assets/images/icons_clientes/nestle.jpg'),
	(57, 'OI', '/assets/images/icons_clientes/oi.jpg'),
	(58, 'PEUGEOT', '/assets/images/icons_clientes/pegeout.jpg'),
	(59, 'PRAVALER', '/assets/images/icons_clientes/pravaler.jpg'),
	(60, 'QUALICORP', '/assets/images/icons_clientes/qualicorp.jpg'),
	(61, 'QUINTO ANDAR', '/assets/images/icons_clientes/quintoandar.jpg'),
	(62, 'RENNER', '/assets/images/icons_clientes/renner.jpg'),
	(63, 'RIOT GAMES', '/assets/images/icons_clientes/riotgames.jpg'),
	(64, 'SAMSUNG', '/assets/images/icons_clientes/samsung.jpg'),
	(65, 'SANTANDER', '/assets/images/icons_clientes/santander.jpg'),
	(66, 'SEM PARAR', '/assets/images/icons_clientes/semparar.jpg'),
	(67, 'SHELL', '/assets/images/icons_clientes/shell.jpg'),
	(68, 'SODEXO', '/assets/images/icons_clientes/sodexo.jpg'),
	(69, 'SONY', '/assets/images/icons_clientes/sony.jpg'),
	(70, 'STELO', '/assets/images/icons_clientes/stelo.jpg'),
	(71, 'SUL AMERICA CIA', '/assets/images/icons_clientes/sulamerica.jpg'),
	(72, 'TELEFONICA', '/assets/images/icons_clientes/telefonica.jpg'),
	(73, 'TIM', '/assets/images/icons_clientes/tim.jpg'),
	(75, 'UNIDAS', '/assets/images/icons_clientes/unidas.jpg'),
	(76, 'UNILEVER', '/assets/images/icons_clientes/unilever.jpg'),
	(77, 'UNIMED BH', '/assets/images/icons_clientes/unimedbh.jpg'),
	(78, 'UNIMED FLORIANOPOLIS', '/assets/images/icons_clientes/unimedfloripa.jpg'),
	(79, 'UNIMED RJ', '/assets/images/icons_clientes/unimedrio.jpg'),
	(80, 'VELOE', '/assets/images/icons_clientes/veloe.jpg'),
	(81, 'VIA VAREJO', '/assets/images/icons_clientes/viavarejo.jpg'),
	(82, 'VIVO', '/assets/images/icons_clientes/vivo.jpg'),
	(83, 'WHITE MARTINS BRASIL', '/assets/images/icons_clientes/whitemartins.jpg'),
	(84, 'XP INVESTIMENTOS', '/assets/images/icons_clientes/xp.jpg'),
	(86, 'Fleury', '/assets/images/icons_clientes/fleury.jpg'),
	(87, 'Carrefour', '/assets/images/icons_clientes/carrefour.jpg'),
	(88, 'Froneri', '/assets/images/icons_clientes/froneri.jpg'),
	(89, 'UNIMED', '/assets/images/icons_clientes/unimed.jpg'),
	(91, 'UNIMED PORTO ALEGRE', '/assets/images/icons_clientes/unimedpoa.jpg'),
	(92, '3M', '/assets/images/icons_clientes/3m.jpg'),
	(93, 'JCA', '/assets/images/icons_clientes/jca.jpg'),
	(94, 'CVC', '/assets/images/icons_clientes/cvc.jpg'),
	(95, 'PASA', '/assets/images/icons_clientes/pasa.jpg'),
	(96, 'FACILY', '/assets/images/icons_clientes/facily.jpg');

INSERT INTO SITES (IDSITE, NOME) VALUES
	(1, 'GERAL'),
	(2, 'BELEM'),
	(3, 'CAMPO GRANDE'),
	(4, 'CASA'),
	(5, 'CIDADE NOVA'),
	(6, 'DEL CASTILHO'),
	(7, 'FEIRA DE SANTANA'),
	(8, 'FORTALEZA'),
	(9, 'GOIANIA'),
	(10, 'GUARULHOS'),
	(11, 'INTERFILE'),
	(12, 'LIBERDADE'),
	(13, 'MADUREIRA'),
	(14, 'NOVA SAO PAULO'),
	(15, 'NOVA SAO PAULO 2'),
	(16, 'OLIVEIRA COUTINHO'),
	(17, 'PRADO'),
	(18, 'PENHA'),
	(19, 'PORTO ALEGRE'),
	(20, 'REPUBLICA'),
	(21, 'RIBEIRAO PRETO'),
	(22, 'ROCHAVERA'),
	(23, 'SALVADOR'),
	(24, 'SANTANA'),
	(25, 'SANTO ANDRE'),
	(26, 'SANTO ANTONIO'),
	(27, 'SANTOS'),
	(28, 'SAO BENTO 1'),
	(29, 'SAO BENTO 2'),
	(30, 'SAO BERNARDO DO CAMPO'),
	(31, 'SAO CAETANO DO SUL'),
	(32, 'SAO CRISTOVÃO'),
	(33, 'SAO JOSE DOS CAMPOS'),
	(34, 'TELEPORTO'),
	(35, 'URUGUAI'),
	(36, 'ZONA SUL'),
	(37, 'ZONA LESTE'),
	(38, 'SEDE');

INSERT INTO DOMINIOS (IDDOMINIO, NOME) VALUES
	(1, 'ATENTOBR'),
	(2, 'ACIELO'),
	(3, 'AFACEBOOKBR'),
	(4, 'APOIOCASASBAHIA'),
	(5, 'ICATU-ATENTO'),
	(6, 'ARENNERBR'),
	(7, 'ITAUATENTOBR'),
	(8, 'AINTERMEDICA'),
	(9, 'CLIENTEAPBR'),
	(10, 'RECOVERYBR');

INSERT INTO DACS (IDDAC, NOME) VALUES
	(1, 'CLIENTE'),
	(2, 'AVAYA CLOUD'),
	(3, 'BLZ A/ 0*25/ 10.189.0.9192'),
	(4, 'BLZ B / 0*28 / 10.189.0.90'),
	(5, 'BLZ C / 0*35 / 10.189.0.93'),
	(6, 'CAB / 0*71 / 10.181.0.90'),
	(7, 'GOI / 0*62 / 10.157.0.90'),
	(8, 'PEN A / 0*20 / 10.191.0.90'),
	(9, 'PEN B / 0*21 / 10.191.0.91'),
	(10, 'POA / 0*51 / 10.7.0.90'),
	(11, 'PRD / 0*31 / 10.141.0.90'),
	(12, 'RP / 0*16 / 10.10.0.90'),
	(13, 'SBC / 0*15 / 10.2.0.90'),
	(14, 'SBE A / 0*17 / 10.155.2.90'),
	(15, 'SBE B / 0*27 / 10.155.0.92'),
	(16, 'SBE C / 0*37'),
	(17, 'SBE D / 0*18 / 10.155.0.93'),
	(18, 'SBE E / 0*19 / 10.155.2.91'),
	(19, 'SJC / 0*12 / 10.4.0.90');

-- --------------------------------------------------------
-- Usuário inicial
-- Login NT: adm
-- Senha: adm
-- --------------------------------------------------------
INSERT INTO USUARIOS (IDUSUARIO, NOME, LOGIN_NT, RE, CARGO, EMAIL, SENHA, V_USUARIOS, V_BDC_POSTS, V_BDC_ADM, V_IMDB, V_GSA, V_MAPA_OPERACIONAL, V_MAPA_OPERACIONAL_ADM, ID_SITE, DATA_CRIACAO) VALUES
	(1, 'ADM', 'adm', '1111111', 'ADM', 'adm@atento.com.br', '$2a$10$qBilPcTghY3kr3p1llu9SeEA3xH4sMkXidTYkAtyfTkCdmO50YlLq', 'S', 'S', 'S', 'S', 'S', 'S', 'S', 24, '2022-03-21 00:28:09');
