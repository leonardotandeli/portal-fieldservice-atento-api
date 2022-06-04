package modelos

// DadosAutenticacao contém o token e o id do usuário autenticado, além de outras informações que ficam armazenadas no cookie
type DadosAD struct {
	LOGIN_NT               string `json:"login_nt"`
	NOME                   string `json:"nome,omitempty"`
	CONTA_ATIVA            string `json:"conta_ativa,omitempty"`
	SENHA_ULTIMA_DEFINICAO string `json:"senha_ultima_definicao,omitempty"`
	SENHA_EXPIRACAO        string `json:"senha_expiracao,omitempty"`
	DATA_ULTIMO_LOGON      string `json:"data_ultimo_logon,omitempty"`
}
