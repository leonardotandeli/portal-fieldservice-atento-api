package modelos

// DAC representa um dac cadastrado em sistema
type Dac struct {
	IDDAC uint64 `json:"iddac,omitempty"`
	NOME  string `json:"nome,omitempty"`
}
