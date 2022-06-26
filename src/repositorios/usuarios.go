package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

// Usuarios representa um repositório de usuarios
type Usuarios struct {
	db *sql.DB
}

// NovoRepositorioDeUsuarios inicia um repositório de usuários
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Criar insere um novo usuário no banco de dados
func (repositorio Usuarios) CriarUsuario(usuario modelos.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"INSERT INTO USUARIOS(NOME, LOGIN_NT, RE, CARGO, EMAIL, SENHA, V_USUARIOS, V_BDC_POSTS, V_BDC_ADM, V_IMDB, V_GSA, V_MAPA_OPERACIONAL, V_MAPA_OPERACIONAL_ADM, ID_SITE) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.NOME, usuario.LOGIN_NT, usuario.RE, usuario.CARGO, usuario.EMAIL, usuario.SENHA, usuario.V_USUARIOS, usuario.V_BDC_POSTS, usuario.V_BDC_ADM, usuario.V_IMDB, usuario.V_GSA, usuario.V_MAPA_OPERACIONAL, usuario.V_MAPA_OPERACIONAL_ADM, usuario.ID_SITE)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil

}

// Buscar traz todos os usuários que atendem um filtro de nome ou login
func (repositorio Usuarios) Buscar(nomeOuLogin string) ([]modelos.Usuario, error) {

	//filtro nome ou login
	nomeOuLogin = fmt.Sprintf("%%%s%%", nomeOuLogin)

	linhas, erro := repositorio.db.Query(
		"SELECT U.IDUSUARIO, U.NOME, U.RE, U.LOGIN_NT, U.CARGO, U.EMAIL, U.V_USUARIOS, U.V_BDC_POSTS, U.V_BDC_ADM, U.V_IMDB, U.V_GSA, U.V_MAPA_OPERACIONAL, U.V_MAPA_OPERACIONAL_ADM, U.ID_SITE, S.IDSITE, S.NOME FROM USUARIOS U INNER JOIN SITES S ON S.IDSITE = U.ID_SITE WHERE U.NOME LIKE ? or U.LOGIN_NT LIKE ? ORDER BY U.IDUSUARIO DESC", nomeOuLogin, nomeOuLogin)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario

		if erro = linhas.Scan(
			&usuario.IDUSUARIO,
			&usuario.NOME,
			&usuario.RE,
			&usuario.LOGIN_NT,
			&usuario.CARGO,
			&usuario.EMAIL,
			&usuario.V_USUARIOS,
			&usuario.V_BDC_POSTS,
			&usuario.V_BDC_ADM,
			&usuario.V_IMDB,
			&usuario.V_GSA,
			&usuario.V_MAPA_OPERACIONAL,
			&usuario.V_MAPA_OPERACIONAL_ADM,
			&usuario.ID_SITE,
			&usuario.Site.IDSITE,
			&usuario.Site.NOME,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

// BuscarPorID traz um usuário do banco de dados filtrado pelo id
func (repositorio Usuarios) BuscarPorID(ID uint64) (modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(
		"SELECT U.IDUSUARIO, U.NOME, U.RE, U.LOGIN_NT, U.CARGO, U.EMAIL, U.V_USUARIOS, U.V_BDC_POSTS, U.V_BDC_ADM, U.V_IMDB, U.V_GSA, U.V_MAPA_OPERACIONAL, U.V_MAPA_OPERACIONAL_ADM, U.ID_SITE, S.IDSITE, S.NOME FROM USUARIOS U INNER JOIN SITES S ON S.IDSITE = U.ID_SITE WHERE U.IDUSUARIO = ?", ID)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linhas.Close()

	var usuario modelos.Usuario

	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.IDUSUARIO,
			&usuario.NOME,
			&usuario.RE,
			&usuario.LOGIN_NT,
			&usuario.CARGO,
			&usuario.EMAIL,
			&usuario.V_USUARIOS,
			&usuario.V_BDC_POSTS,
			&usuario.V_BDC_ADM,
			&usuario.V_IMDB,
			&usuario.V_GSA,
			&usuario.V_MAPA_OPERACIONAL,
			&usuario.V_MAPA_OPERACIONAL_ADM,
			&usuario.ID_SITE,
			&usuario.Site.IDSITE,
			&usuario.Site.NOME,
		); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil
}

// Atualizar altera as informações de um usuário no banco de dados
func (repositorio Usuarios) Atualizar(ID uint64, usuario modelos.Usuario) error {
	statement, erro := repositorio.db.Prepare(
		"UPDATE USUARIOS SET NOME = ?, RE = ?, LOGIN_NT = ?, CARGO = ?, EMAIL = ?, V_USUARIOS = ?, V_BDC_POSTS = ?, V_BDC_ADM = ?, V_IMDB = ?, V_GSA = ?, V_MAPA_OPERACIONAL = ?, V_MAPA_OPERACIONAL_ADM = ?, ID_SITE = ? WHERE IDUSUARIO = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuario.NOME, usuario.RE, usuario.LOGIN_NT, usuario.CARGO, usuario.EMAIL, usuario.V_USUARIOS, usuario.V_BDC_POSTS, usuario.V_BDC_ADM, usuario.V_IMDB, usuario.V_GSA, usuario.V_MAPA_OPERACIONAL, usuario.V_MAPA_OPERACIONAL_ADM, usuario.ID_SITE, ID); erro != nil {
		return erro
	}

	return nil
}

// Deletar exclui o usuário do banco de dados
func (repositorio Usuarios) Deletar(ID uint64) error {
	statement, erro := repositorio.db.Prepare("DELETE FROM USUARIOS WHERE IDUSUARIO = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

// BuscarPorEmail busca um usuário por email e retorna o seu id e senha com hash
func (repositorio Usuarios) BuscarPorEmail(email string) (modelos.Usuario, error) {
	linha, erro := repositorio.db.Query("SELECT U.IDUSUARIO, U.LOGIN_NT, U.NOME, U.CARGO, U.EMAIL, U.V_USUARIOS, U.V_BDC_POSTS, U.V_BDC_ADM, U.V_IMDB, U.V_GSA, U.V_MAPA_OPERACIONAL, U.V_MAPA_OPERACIONAL_ADM, S.IDSITE, S.NOME, U.SENHA FROM USUARIOS U INNER JOIN SITES S ON S.IDSITE = U.ID_SITE = U.IDUSUARIO WHERE EMAIL = ?", email)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linha.Close()

	var usuario modelos.Usuario

	if linha.Next() {
		if erro = linha.Scan(&usuario.IDUSUARIO, &usuario.LOGIN_NT, &usuario.NOME, &usuario.CARGO, &usuario.EMAIL, &usuario.V_USUARIOS, &usuario.V_BDC_POSTS, &usuario.V_BDC_ADM, &usuario.V_IMDB, &usuario.V_GSA, &usuario.V_MAPA_OPERACIONAL, &usuario.V_MAPA_OPERACIONAL_ADM, &usuario.Site.IDSITE, &usuario.Site.NOME, &usuario.SENHA); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil

}

// BuscarPorLogin busca um usuário pelo Login
func (repositorio Usuarios) BuscarPorLogin(LOGIN_NT string) (modelos.Usuario, error) {
	linha, erro := repositorio.db.Query("SELECT U.IDUSUARIO, U.LOGIN_NT, U.RE, U.NOME, U.CARGO, U.EMAIL, U.V_USUARIOS, U.V_BDC_POSTS, U.V_BDC_ADM, U.V_IMDB, U.V_GSA, U.V_MAPA_OPERACIONAL, U.V_MAPA_OPERACIONAL_ADM, S.IDSITE, S.NOME, U.SENHA FROM USUARIOS U INNER JOIN SITES S ON S.IDSITE = U.ID_SITE WHERE U.LOGIN_NT = ?", LOGIN_NT)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linha.Close()

	var usuario modelos.Usuario

	if linha.Next() {
		if erro = linha.Scan(&usuario.IDUSUARIO, &usuario.LOGIN_NT, &usuario.RE, &usuario.NOME, &usuario.CARGO, &usuario.EMAIL, &usuario.V_USUARIOS, &usuario.V_BDC_POSTS, &usuario.V_BDC_ADM, &usuario.V_IMDB, &usuario.V_GSA, &usuario.V_MAPA_OPERACIONAL, &usuario.V_MAPA_OPERACIONAL_ADM, &usuario.Site.IDSITE, &usuario.Site.NOME, &usuario.SENHA); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil

}

// BuscarSenha traz a senha de um usuário pelo ID
func (repositorio Usuarios) BuscarSenha(usuarioID uint64) (string, error) {
	linha, erro := repositorio.db.Query("SELECT SENHA FROM USUARIOS WHERE IDUSUARIO = ?", usuarioID)
	if erro != nil {
		return "", erro
	}
	defer linha.Close()

	var usuario modelos.Usuario

	if linha.Next() {
		if erro = linha.Scan(&usuario.SENHA); erro != nil {
			return "", erro
		}
	}

	return usuario.SENHA, nil
}

// AtualizarSenha altera a senha de um usuário
func (repositorio Usuarios) AtualizarSenha(usuarioID uint64, senha string) error {
	statement, erro := repositorio.db.Prepare("UPDATE USUARIOS SET SENHA = ? WHERE IDUSUARIO = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	fmt.Println(usuarioID)
	if _, erro = statement.Exec(senha, usuarioID); erro != nil {
		return erro
	}

	return nil
}
