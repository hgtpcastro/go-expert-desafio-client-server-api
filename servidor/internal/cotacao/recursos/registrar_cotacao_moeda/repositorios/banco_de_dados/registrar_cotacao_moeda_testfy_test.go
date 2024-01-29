package bancodedados

import (
	"context"
	"database/sql"
	"log"
	"testing"
	"time"

	"github.com/hgtpcastro/go-expert/desafio-client-server-api/servidor/internal/cotacao/entidade"
	"github.com/hgtpcastro/go-expert/desafio-client-server-api/servidor/internal/cotacao/recursos/registrar_cotacao_moeda/contratos"
	"github.com/stretchr/testify/suite"

	// sqlite3
	_ "github.com/mattn/go-sqlite3"
)

type RepositorioRegistrarCotacaoMoedaTestSuite struct {
	suite.Suite
	bancoDados  *sql.DB
	repositorio contratos.RepositorioCotacaoMoeda
}

func (suiteTeste *RepositorioRegistrarCotacaoMoedaTestSuite) SetupSuite() {
	log.Println("SetupSuite...")
	bancoDadosTeste, erro := sql.Open("sqlite3", ":memory:")
	suiteTeste.NoError(erro)
	bancoDadosTeste.Exec(`
		CREATE TABLE IF NOT EXISTS cotacao_moeda (
			id varchar(255) NOT NULL PRIMARY KEY,
			moeda varchar(255),
			de varchar(255),
			para varchar(255),
			nome varchar(255),			
			valor varchar(255),						
			data varchar(255)
		);
	`)
	suiteTeste.bancoDados = bancoDadosTeste

	suiteTeste.repositorio = NovoRepositorioRegistrarCotacaoMoeda(suiteTeste.bancoDados)
	suiteTeste.NotNil(suiteTeste.repositorio)
}

func (suiteTeste *RepositorioRegistrarCotacaoMoedaTestSuite) TearDownTest() {
	log.Println("TearDownTest...")
	suiteTeste.bancoDados.Close()
}

func TestSuite(t *testing.T) {
	log.Println("TestSuite...")
	suite.Run(t, new(RepositorioRegistrarCotacaoMoedaTestSuite))
}

func (suiteTeste *RepositorioRegistrarCotacaoMoedaTestSuite) RegistrarCotacaoMoedaNoBancoDeDados() {
	cotacaoMoeda := entidade.NovoCotacaoMoeda(
		"USD-BRL",
		"USD",
		"BRL",
		"Dólar Americano/Real Brasileiro",
		"4.9168",
		"2024-01-26 18:59:33",
	)
	suiteTeste.Nil(cotacaoMoeda)

	// repositorio := NovoRepositorioRegistrarCotacaoMoeda(suiteTeste.bancoDados)
	// suiteTeste.NotNil(repositorio)

	ctx, cancelar := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancelar()

	erro := suiteTeste.repositorio.Registrar(ctx, cotacaoMoeda)
	suiteTeste.NoError(erro)

	var cotacaoMoedaEsperado entidade.CotacaoMoeda

	erro = suiteTeste.bancoDados.
		QueryRow(`
			Select 
				id, 
				moeda, 
				de, 
				para, 
				nome 
			from cotacao_moeda 
			where id = ?`,
			cotacaoMoeda.Id,
		).
		Scan(
			&cotacaoMoedaEsperado.Id,
			&cotacaoMoedaEsperado.Moeda,
			&cotacaoMoedaEsperado.De,
			&cotacaoMoedaEsperado.Para,
			&cotacaoMoedaEsperado.Valor,
		)
	suiteTeste.NoError(erro)
	suiteTeste.Equal(cotacaoMoeda.Id, cotacaoMoedaEsperado.Id)
}
