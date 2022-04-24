package modelos

import (
	"time"
)

// Dominio representa os domínios cadastrados no sistema
type Dominio struct {
	IDDOMINIO uint64 `json:"iddominio,omitempty"`
	NOME      string `json:"nome,omitempty"`
}

// DAC representa um dac cadastrado em sistema
type Dac struct {
	IDDAC uint64 `json:"iddac,omitempty"`
	NOME  string `json:"nome,omitempty"`
}

// Cliente representa um cliente cadastrado em sistema
type Cliente struct {
	IDCLIENTE uint64 `json:"idcliente,omitempty"`
	NOME      string `json:"nome,omitempty"`
	LOGO_URL  string `json:"logo_url,omitempty"`
}

// Site representa os sites cadastrados no sistema
type Site struct {
	IDSITE     uint64 `json:"idsite,omitempty"`
	NOME       string `json:"nome,omitempty"`
	ID_USUARIO string `json:"id_usuario,omitempty"`
}

// MapaOperacional representa as operações cadastrados no sistema
type MapaOperacional struct {
	IDMAPA            uint64    `json:"idmapa,omitempty"`
	OPERACAO          string    `json:"operacao,omitempty"`
	VLAN_DADOS        string    `json:"vlan_dados,omitempty"`
	VLAN_VOZ          string    `json:"vlan_voz,omitempty"`
	CONFIG_CONTRATUAL string    `json:"config_contratual,omitempty"`
	VERSAO_WINDOWS    string    `json:"versao_windows,omitempty"`
	IMAGEM            string    `json:"imagem,omitempty"`
	TEMPLATE          string    `json:"template,omitempty"`
	GRUPO_IMDB        string    `json:"grupo_imdb,omitempty"`
	GRAVADOR          string    `json:"gravador,omitempty"`
	OBSERVACOES       string    `json:"observacoes,omitempty"`
	ID_SITE           string    `json:"id_site,omitempty"`
	ID_CLIENTE        string    `json:"id_cliente,omitempty"`
	ID_DOMINIO        string    `json:"id_dominio,omitempty"`
	ID_DAC            string    `json:"id_dac,omitempty"`
	DATA_CRIACAO      time.Time `json:"data_criacao,omitempty"`
	Site              Site
	Cliente           Cliente
	Dominio           Dominio
	Dac               Dac
}
