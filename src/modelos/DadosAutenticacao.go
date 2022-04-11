package modelos

// DadosAutenticacao contém o token e o id do usuário autenticado, além de outras informações que ficam armazenadas no cookie
type DadosAutenticacao struct {
	ID                 string `json:"id"`
	Token              string `json:"token"`
	NOME               string `json:"nome,omitempty"`
	LOGIN_NT           string `json:"login_nt,omitempty"`
	RE                 string `json:"re,omitempty"`
	CARGO              string `json:"cargo,omitempty"`
	EMAIL              string `json:"email,omitempty"`
	SENHA              string `json:"senha,omitempty"`
	V_USUARIOS         string `json:"v_usuarios,omitempty"`
	V_BDC_POSTS        string `json:"v_bdc_posts,omitempty"`
	V_BDC_ADM          string `json:"v_bdc_adm,omitempty"`
	V_IMDB             string `json:"v_imdb,omitempty"`
	V_GSA              string `json:"v_gsa,omitempty"`
	V_MAPA_OPERACIONAL string `json:"v_mapa_operacional,omitempty"`
	Site               Site
}
