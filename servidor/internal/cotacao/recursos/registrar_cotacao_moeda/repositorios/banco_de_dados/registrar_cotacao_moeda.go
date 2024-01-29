package bancodedados

import (
	"context"
	"database/sql"
	"log"

	"github.com/hgtpcastro/go-expert/desafio-client-server-api/servidor/internal/cotacao/entidade"
)

type RepositorioRegistrarCotacaoMoeda struct {
	bancoDados *sql.DB
}

func NovoRepositorioRegistrarCotacaoMoeda(bancoDados *sql.DB) *RepositorioRegistrarCotacaoMoeda {
	r := &RepositorioRegistrarCotacaoMoeda{
		bancoDados: bancoDados,
	}
	r.CriarTabelaCotacaoMoeda()
	return r
}

func (r *RepositorioRegistrarCotacaoMoeda) CriarTabelaCotacaoMoeda() error {
	_, erro := r.bancoDados.Exec(`
		CREATE TABLE IF NOT EXISTS cotacao (
			id varchar(255) NOT NULL PRIMARY KEY,
			moeda varchar(255),
			de varchar(255),
			para varchar(255),
			nome varchar(255),			
			valor varchar(255),						
			data varchar(255)
		);
	`)
	if erro != nil {
		return erro
	}

	return nil
}

func (r *RepositorioRegistrarCotacaoMoeda) Registrar(contexto context.Context, cotacaoMoeda *entidade.CotacaoMoeda) error {
	stmt, erro := r.bancoDados.PrepareContext(contexto,
		` insert into cotacao(id, moeda, de, para, nome, valor, data)
			values(?, ?, ?, ?, ?, ?, ?)
		`,
	)
	if erro != nil {
		log.Println(erro)
		return erro
	}

	// time.Sleep(11 * time.Millisecond)

	_, erro = stmt.ExecContext(contexto,
		cotacaoMoeda.Id,
		cotacaoMoeda.Moeda,
		cotacaoMoeda.De,
		cotacaoMoeda.Para,
		cotacaoMoeda.Nome,
		cotacaoMoeda.Valor,
		cotacaoMoeda.Data,
	)
	if erro != nil {
		log.Println(erro)
		return erro
	}

	return nil
}
