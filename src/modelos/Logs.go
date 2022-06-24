package modelos

import "time"

type Logs struct {
	Usuario Usuario
	ACTION  string    `json:"action,omitempty"`
	DATA    time.Time `json:"data,omitempty"`
}
