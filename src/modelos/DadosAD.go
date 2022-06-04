package modelos

// DadosAutenticacao contém o token e o id do usuário autenticado, além de outras informações que ficam armazenadas no cookie
type DadosAD struct {
	LOGIN_NT               string `json:"login_nt"`
	NOME                   string `json:"nome,omitempty"`
	OPERACAO               string `json:"operacao,omitempty"`
	CONTA_ATIVA            string `json:"conta_ativa,omitempty"`
	CONTA_EXPIRA_EM        string `json:"conta_expira_em,omitempty"`
	SENHA_ULTIMA_DEFINICAO string `json:"senha_ultima_definicao,omitempty"`
	SENHA_EXPIRACAO        string `json:"senha_expiracao,omitempty"`
	SENHA_ALTERACAO        string `json:"senha_alteracao,omitempty"`
	DATA_ULTIMO_LOGON      string `json:"data_ultimo_logon,omitempty"`
	GPOS                   string `json:"gpos,omitempty"`
}
