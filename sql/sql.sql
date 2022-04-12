DROP DATABASE IF EXISTS KEEPER; /*NOME DO BANCO*/
CREATE DATABASE IF NOT EXISTS KEEPER;
USE KEEPER;

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
  CARGO ENUM('ADM', 'M3', 'M2', 'M1', 'P4', 'P3', 'P2', 'P1', 'N2') NOT NULL,
  EMAIL VARCHAR(50) UNIQUE KEY NOT NULL,
  SENHA VARCHAR(100) NOT NULL,
  V_USUARIOS ENUM('S','N') NOT NULL,
  V_BDC_POSTS ENUM('S','N') NOT NULL,
  V_BDC_ADM ENUM('S','N') NOT NULL,
  V_IMDB ENUM('S','N') NOT NULL,
  V_GSA ENUM('S','N') NOT NULL,
  V_MAPA_OPERACIONAL ENUM('S','N') NOT NULL,
  ID_SITE INT,
  DATA_CRIACAO TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP()
);

ALTER TABLE USUARIOS
ADD V_MAPA_OPERACIONAL_ADM ENUM('S','N') NOT NULL,


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
    NOME VARCHAR(100) NOT NULL
);

CREATE TABLE DOMINIOS(
    IDDOMINIO INT AUTO_INCREMENT PRIMARY KEY,
    NOME VARCHAR(100) NOT NULL
);

CREATE TABLE DACS(
    IDDAC INT AUTO_INCREMENT PRIMARY KEY,
    NOME VARCHAR(100) NOT NULL
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

