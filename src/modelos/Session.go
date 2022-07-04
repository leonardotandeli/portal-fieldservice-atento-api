package modelos

import "time"

//Session representa o registro de um sessão logada.
type Session struct {
	ID                uint64 `json:"id,omitempty"`
	Usuario           Usuario
	DadosAutenticacao DadosAutenticacao
	DATA              time.Time `json:"data,omitempty"`
}
