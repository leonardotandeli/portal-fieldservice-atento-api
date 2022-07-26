package modelos

import "time"

//Session representa o registro de um sessão logada.
type Session struct {
	ID         uint64    `json:"id,omitempty" gorm:"primaryKey"`
	ID_USUARIO uint64    `json:"nome,omitempty"  gorm:"column:ID_USUARIO"`
	Token      string    `json:"token"`
	DATA       time.Time `json:"data,omitempty"`
}
