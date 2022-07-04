package modelos

import (
	"api/src/seguranca"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// Usuario representa um usuário cadastrado no sistema
type Usuario struct {
	IDUSUARIO              uint64    `json:"idusuario,omitempty"`
	NOME                   string    `json:"nome,omitempty""`
	LOGIN_NT               string    `json:"login_nt,omitempty"`
	RE                     string    `json:"re,omitempty"`
	CARGO                  string    `json:"cargo,omitempty"`
	EMAIL                  string    `json:"email,omitempty"`
	SENHA                  string    `json:"senha,omitempty"`
	V_USUARIOS             string    `json:"v_usuarios,omitempty"`
	V_BDC_POSTS            string    `json:"v_bdc_posts,omitempty"`
	V_BDC_ADM              string    `json:"v_bdc_adm,omitempty"`
	V_IMDB                 string    `json:"v_imdb,omitempty"`
	V_GSA                  string    `json:"v_gsa,omitempty"`
	V_MAPA_OPERACIONAL     string    `json:"v_mapa_operacional,omitempty"`
	V_MAPA_OPERACIONAL_ADM string    `json:"v_mapa_operacional_adm,omitempty"`
	ID_SITE                string    `json:"id_site,omitempty"`
	DATA_CRIACAO           time.Time `json:"data_criacao,omitempty"`
	Site                   Site
	STATUS                 string `json:"status,omitempty"`
}

// Struct Senha representa o formato da requisição de alteração de senha
type Senha struct {
	Nova  string `json:"nova"`
	Atual string `json:"atual"`
}

func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}
	if erro := usuario.formatar(etapa); erro != nil {
		return erro
	}

	return nil
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.NOME == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco")
	}

	if usuario.LOGIN_NT == "" {
		return errors.New("O LoginNT é obrigatório e não pode estar em branco")
	}

	if usuario.EMAIL == "" {
		return errors.New("O e-mail é obrigatório e não pode estar em branco")
	}

	if erro := checkmail.ValidateFormat(usuario.EMAIL); erro != nil {
		return errors.New("O e-mail inserido é inválido")
	}

	if etapa == "cadastro" && usuario.SENHA == "" {
		return errors.New("A senha é obrigatória e não pode estar em branco")
	}

	return nil
}

func (usuario *Usuario) formatar(etapa string) error {
	usuario.NOME = strings.TrimSpace(usuario.NOME)
	usuario.EMAIL = strings.TrimSpace(usuario.EMAIL)

	if etapa == "cadastro" {
		senhaComHash, erro := seguranca.Hash(usuario.SENHA)
		if erro != nil {
			return erro
		}

		usuario.SENHA = string(senhaComHash)
	}

	return nil
}
