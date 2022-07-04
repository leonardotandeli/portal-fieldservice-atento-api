package modelos

import (
	"time"
)

// MapaOperacional representa uma operação cadastrada no sistema
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
