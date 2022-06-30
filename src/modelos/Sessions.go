package modelos

import "time"

type Session struct {
	ID                uint64 `json:"id,omitempty"`
	Usuario           Usuario
	DadosAutenticacao DadosAutenticacao
	DATA              time.Time `json:"data,omitempty"`
}
