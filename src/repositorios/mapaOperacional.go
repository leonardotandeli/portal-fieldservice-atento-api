package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
	"math"
	"strconv"
)

type MapasOperacional struct {
	db *sql.DB
}

func NovoRepositorioDeMapasOperacional(db *sql.DB) *MapasOperacional {
	return &MapasOperacional{db}
}

func (repositorio MapasOperacional) CriarDadosMapa(mapa modelos.MapaOperacional) (uint64, error) {

	statement, erro := repositorio.db.Prepare(
		"INSERT INTO MAPA_OPERACIONAL(OPERACAO, VLAN_DADOS, VLAN_VOZ, CONFIG_CONTRATUAL, VERSAO_WINDOWS, IMAGEM, TEMPLATE, GRUPO_IMDB, GRAVADOR, OBSERVACOES, ID_SITE, ID_CLIENTE, ID_DOMINIO, ID_DAC) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(mapa.OPERACAO, mapa.VLAN_DADOS, mapa.VLAN_VOZ, mapa.CONFIG_CONTRATUAL, mapa.VERSAO_WINDOWS, mapa.IMAGEM, mapa.TEMPLATE, mapa.GRUPO_IMDB, mapa.GRAVADOR, mapa.OBSERVACOES, mapa.ID_SITE, mapa.ID_CLIENTE, mapa.ID_DOMINIO, mapa.ID_DAC)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil

}

// Buscar traz todas as entradas do mapa operacional
func (repositorio MapasOperacional) Buscar(urlSite string, urlCliente string, urlPage string) ([]modelos.MapaOperacional, error) {

	urlSite = fmt.Sprintf("%s", urlSite)       // filtra por parametro na url ?site=[id]
	urlCliente = fmt.Sprintf("%s", urlCliente) // filtra por parametro na url ?site=[id]

	urlPage = fmt.Sprintf("%s", urlPage) // filtra por parametro na url ?site=[id]

	if urlPage == "" {
		urlPage = "1"
	}

	urlPageInt, _ := strconv.Atoi(urlPage)
	porPagina := 9

	var total int
_:
	repositorio.db.QueryRow("SELECT COUNT(IDMAPA) FROM mapa_operacional").Scan(&total)

	fmt.Println(total)

	linhas, erro := repositorio.db.Query(
		"SELECT M.IDMAPA, M.OPERACAO, M.VLAN_DADOS, M.VLAN_VOZ, M.CONFIG_CONTRATUAL, M.VERSAO_WINDOWS, M.IMAGEM, M.TEMPLATE, M.GRUPO_IMDB, M.GRAVADOR, M.OBSERVACOES, M.ID_SITE, M.ID_CLIENTE, M.ID_DOMINIO, M.ID_DAC, S.NOME, C.NOME, D.NOME, T.NOME FROM MAPA_OPERACIONAL M INNER JOIN SITES S ON M.ID_SITE = S.IDSITE INNER JOIN CLIENTES C ON M.ID_CLIENTE = C.IDCLIENTE INNER JOIN DOMINIOS D ON M.ID_DOMINIO = D.IDDOMINIO INNER JOIN DACS T ON M.ID_DAC = T.IDDAC ORDER BY M.IDMAPA DESC LIMIT ? OFFSET ?", porPagina, (urlPageInt-1)*porPagina)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var MapasOperacional []modelos.MapaOperacional

	for linhas.Next() {
		var mapa_operacional modelos.MapaOperacional
		mapa_operacional.Pagination.Total = total
		mapa_operacional.Pagination.Pagina = urlPageInt
		mapa_operacional.Pagination.UltimaPagina = math.Ceil(float64(total) / float64(porPagina))

		if erro = linhas.Scan(
			&mapa_operacional.IDMAPA,
			&mapa_operacional.OPERACAO,
			&mapa_operacional.VLAN_DADOS,
			&mapa_operacional.VLAN_VOZ,
			&mapa_operacional.CONFIG_CONTRATUAL,
			&mapa_operacional.VERSAO_WINDOWS,
			&mapa_operacional.IMAGEM,
			&mapa_operacional.TEMPLATE,
			&mapa_operacional.GRUPO_IMDB,
			&mapa_operacional.GRAVADOR,
			&mapa_operacional.OBSERVACOES,
			&mapa_operacional.ID_SITE,
			&mapa_operacional.ID_CLIENTE,
			&mapa_operacional.ID_DOMINIO,
			&mapa_operacional.ID_DAC,
			&mapa_operacional.Site.NOME,
			&mapa_operacional.Cliente.NOME,
			&mapa_operacional.Dominio.NOME,
			&mapa_operacional.Dac.NOME,
		); erro != nil {
			return nil, erro
		}

		MapasOperacional = append(MapasOperacional, mapa_operacional)
	}

	return MapasOperacional, nil

}

var stringC string

//  BuscarString traz todas as entradas do mapa operacional filtrando pelo site ou cliente
func (repositorio MapasOperacional) BuscarString(urlSite string, urlCliente string) ([]modelos.MapaOperacional, error) {

	urlSite = fmt.Sprintf("%s", urlSite)       // filtra por parametro na url ?site=[id]
	urlCliente = fmt.Sprintf("%s", urlCliente) // filtra por parametro na url ?site=[id]

	if urlSite == "" && urlCliente != "" {
		stringC = "SELECT M.IDMAPA, M.OPERACAO, M.VLAN_DADOS, M.VLAN_VOZ, M.CONFIG_CONTRATUAL, M.VERSAO_WINDOWS, M.IMAGEM, M.TEMPLATE, M.GRUPO_IMDB, M.GRAVADOR, M.OBSERVACOES, M.ID_SITE, M.ID_CLIENTE, M.ID_DOMINIO, M.ID_DAC, S.NOME, C.NOME, D.NOME, T.NOME FROM MAPA_OPERACIONAL M INNER JOIN SITES S ON M.ID_SITE = S.IDSITE INNER JOIN CLIENTES C ON M.ID_CLIENTE = C.IDCLIENTE INNER JOIN DOMINIOS D ON M.ID_DOMINIO = D.IDDOMINIO INNER JOIN DACS T ON M.ID_DAC = T.IDDAC WHERE S.IDSITE = ? OR C.IDCLIENTE = ? ORDER BY M.IDMAPA DESC"
	} else if urlSite != "" && urlCliente == "" {
		stringC = "SELECT M.IDMAPA, M.OPERACAO, M.VLAN_DADOS, M.VLAN_VOZ, M.CONFIG_CONTRATUAL, M.VERSAO_WINDOWS, M.IMAGEM, M.TEMPLATE, M.GRUPO_IMDB, M.GRAVADOR, M.OBSERVACOES, M.ID_SITE, M.ID_CLIENTE, M.ID_DOMINIO, M.ID_DAC, S.NOME, C.NOME, D.NOME, T.NOME FROM MAPA_OPERACIONAL M INNER JOIN SITES S ON M.ID_SITE = S.IDSITE INNER JOIN CLIENTES C ON M.ID_CLIENTE = C.IDCLIENTE INNER JOIN DOMINIOS D ON M.ID_DOMINIO = D.IDDOMINIO INNER JOIN DACS T ON M.ID_DAC = T.IDDAC WHERE S.IDSITE = ? OR C.IDCLIENTE = ? ORDER BY M.IDMAPA DESC"
	} else if urlSite != "" && urlCliente != "" {
		stringC = "SELECT M.IDMAPA, M.OPERACAO, M.VLAN_DADOS, M.VLAN_VOZ, M.CONFIG_CONTRATUAL, M.VERSAO_WINDOWS, M.IMAGEM, M.TEMPLATE, M.GRUPO_IMDB, M.GRAVADOR, M.OBSERVACOES, M.ID_SITE, M.ID_CLIENTE, M.ID_DOMINIO, M.ID_DAC, S.NOME, C.NOME, D.NOME, T.NOME FROM MAPA_OPERACIONAL M INNER JOIN SITES S ON M.ID_SITE = S.IDSITE INNER JOIN CLIENTES C ON M.ID_CLIENTE = C.IDCLIENTE INNER JOIN DOMINIOS D ON M.ID_DOMINIO = D.IDDOMINIO INNER JOIN DACS T ON M.ID_DAC = T.IDDAC WHERE S.IDSITE = ? AND C.IDCLIENTE = ? ORDER BY M.IDMAPA DESC"
	} else {
		stringC = "SELECT M.IDMAPA, M.OPERACAO, M.VLAN_DADOS, M.VLAN_VOZ, M.CONFIG_CONTRATUAL, M.VERSAO_WINDOWS, M.IMAGEM, M.TEMPLATE, M.GRUPO_IMDB, M.GRAVADOR, M.OBSERVACOES, M.ID_SITE, M.ID_CLIENTE, M.ID_DOMINIO, M.ID_DAC, S.NOME, C.NOME, D.NOME, T.NOME FROM MAPA_OPERACIONAL M INNER JOIN SITES S ON M.ID_SITE = S.IDSITE INNER JOIN CLIENTES C ON M.ID_CLIENTE = C.IDCLIENTE INNER JOIN DOMINIOS D ON M.ID_DOMINIO = D.IDDOMINIO INNER JOIN DACS T ON M.ID_DAC = T.IDDAC WHERE S.IDSITE LIKE ? OR C.IDCLIENTE LIKE ? ORDER BY M.IDMAPA DESC"
	}
	//	fmt.Println(stringC)
	linhas, erro := repositorio.db.Query(stringC, urlSite, urlCliente)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var MapasOperacional []modelos.MapaOperacional

	for linhas.Next() {
		var mapa_operacional modelos.MapaOperacional

		if erro = linhas.Scan(
			&mapa_operacional.IDMAPA,
			&mapa_operacional.OPERACAO,
			&mapa_operacional.VLAN_DADOS,
			&mapa_operacional.VLAN_VOZ,
			&mapa_operacional.CONFIG_CONTRATUAL,
			&mapa_operacional.VERSAO_WINDOWS,
			&mapa_operacional.IMAGEM,
			&mapa_operacional.TEMPLATE,
			&mapa_operacional.GRUPO_IMDB,
			&mapa_operacional.GRAVADOR,
			&mapa_operacional.OBSERVACOES,
			&mapa_operacional.ID_SITE,
			&mapa_operacional.ID_CLIENTE,
			&mapa_operacional.ID_DOMINIO,
			&mapa_operacional.ID_DAC,
			&mapa_operacional.Site.NOME,
			&mapa_operacional.Cliente.NOME,
			&mapa_operacional.Dominio.NOME,
			&mapa_operacional.Dac.NOME,
		); erro != nil {
			return nil, erro
		}

		MapasOperacional = append(MapasOperacional, mapa_operacional)
	}

	return MapasOperacional, nil

}

// BuscarPorID busca uma operacao no banco de dados pelo id
func (repositorio MapasOperacional) BuscarPorID(ID uint64) (modelos.MapaOperacional, error) {

	linhas, erro := repositorio.db.Query(
		"SELECT M.IDMAPA, M.OPERACAO, M.VLAN_DADOS, M.VLAN_VOZ, M.CONFIG_CONTRATUAL, M.VERSAO_WINDOWS, M.IMAGEM, M.TEMPLATE, M.GRUPO_IMDB, M.GRAVADOR, M.OBSERVACOES, M.ID_SITE, M.ID_CLIENTE, M.ID_DOMINIO, M.ID_DAC, S.NOME, C.NOME, D.NOME, T.NOME FROM MAPA_OPERACIONAL M INNER JOIN SITES S ON M.ID_SITE = S.IDSITE INNER JOIN CLIENTES C ON M.ID_CLIENTE = C.IDCLIENTE INNER JOIN DOMINIOS D ON M.ID_DOMINIO = D.IDDOMINIO INNER JOIN DACS T ON M.ID_DAC = T.IDDAC WHERE M.IDMAPA = ?", ID)
	if erro != nil {
		return modelos.MapaOperacional{}, erro
	}
	defer linhas.Close()

	var mapa_operacional modelos.MapaOperacional

	if linhas.Next() {
		if erro = linhas.Scan(
			&mapa_operacional.IDMAPA,
			&mapa_operacional.OPERACAO,
			&mapa_operacional.VLAN_DADOS,
			&mapa_operacional.VLAN_VOZ,
			&mapa_operacional.CONFIG_CONTRATUAL,
			&mapa_operacional.VERSAO_WINDOWS,
			&mapa_operacional.IMAGEM,
			&mapa_operacional.TEMPLATE,
			&mapa_operacional.GRUPO_IMDB,
			&mapa_operacional.GRAVADOR,
			&mapa_operacional.OBSERVACOES,
			&mapa_operacional.ID_SITE,
			&mapa_operacional.ID_CLIENTE,
			&mapa_operacional.ID_DOMINIO,
			&mapa_operacional.ID_DAC,
			&mapa_operacional.Site.NOME,
			&mapa_operacional.Cliente.NOME,
			&mapa_operacional.Dominio.NOME,
			&mapa_operacional.Dac.NOME,
		); erro != nil {
			return modelos.MapaOperacional{}, erro
		}
	}

	return mapa_operacional, nil

}

// Atualizar atualizar informações de uma operação no banco de dados utilizando seu ID
func (repositorio MapasOperacional) Atualizar(mapaID uint64, mapa modelos.MapaOperacional) error {
	statement, erro := repositorio.db.Prepare("UPDATE MAPA_OPERACIONAL SET OPERACAO = ?, VLAN_DADOS = ?, VLAN_VOZ = ?, CONFIG_CONTRATUAL = ?, VERSAO_WINDOWS = ?, IMAGEM = ?, TEMPLATE = ?, GRUPO_IMDB = ?, GRAVADOR = ?, OBSERVACOES = ?, ID_SITE = ?, ID_CLIENTE = ?, ID_DOMINIO = ?, ID_DAC = ? WHERE IDMAPA = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(mapa.OPERACAO, mapa.VLAN_DADOS, mapa.VLAN_VOZ, mapa.CONFIG_CONTRATUAL, mapa.VERSAO_WINDOWS, mapa.IMAGEM, mapa.TEMPLATE, mapa.GRUPO_IMDB, mapa.GRAVADOR, mapa.OBSERVACOES, mapa.ID_SITE, mapa.ID_CLIENTE, mapa.ID_DOMINIO, mapa.ID_DAC, mapaID); erro != nil {
		return erro
	}

	return nil
}

// Deletar vai deletar operações pelo ID
func (repositorio MapasOperacional) Deletar(mapaID uint64) error {
	statement, erro := repositorio.db.Prepare("DELETE FROM MAPA_OPERACIONAL WHERE IDMAPA = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(mapaID); erro != nil {
		return erro
	}

	return nil
}
