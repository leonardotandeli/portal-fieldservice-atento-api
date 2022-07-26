package modelos

import "time"

//Logs representa um log cadastrado no sistema
type Logs struct {
	IDUSUARIO uint64    `json:"idusuario,omitempty" gorm:"column:IDUSUARIO"`
	NOME      string    `json:"nome,omitempty"  gorm:"column:NOME"`
	LOGIN_NT  string    `json:"login_nt,omitempty"  gorm:"column:LOGIN_NT"`
	ACTION    string    `json:"action,omitempty"`
	DATA      time.Time `json:"data,omitempty"`
}
