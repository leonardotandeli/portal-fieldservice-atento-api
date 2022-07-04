package modelos

import "time"

//Logs representa um log cadastrado no sistema
type Logs struct {
	Usuario Usuario
	ACTION  string    `json:"action,omitempty"`
	DATA    time.Time `json:"data,omitempty"`
}
