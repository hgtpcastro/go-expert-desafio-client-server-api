package bancodedados

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/hgtpcastro/go-expert/desafio-client-server-api/servidor/internal/cotacao/entidade"

	// sqlite3
	_ "github.com/mattn/go-sqlite3"
)

func TestRegistrarCotacaoMoedaNoBancoDeDados(t *testing.T) {
	t.Run("deve regostrar a cotação no banco de dados", func(t *testing.T) {
		cotacaoMoeda := entidade.NovoCotacaoMoeda(
			"USD-BRL",
			"USD",
			"BRL",
			"Dólar Americano/Real Brasileiro",
			"4.9168",
			"2024-01-26 18:59:33",
		)
		if cotacaoMoeda == nil {
			t.Fatal("esperava que a cotação não fosse nil")
		}

		bancoDados, erro := obterBancoDadosTeste(t)
		if erro != nil {
			t.Fatal(erro)
		}

		repositorio := NovoRepositorioRegistrarCotacaoMoeda(bancoDados)
		if repositorio == nil {
			t.Fatal("esperava que o repositorio não fosse nil")
		}

		ctx, cancelar := context.WithTimeout(context.Background(), 10*time.Millisecond)
		defer cancelar()

		erro = repositorio.Registrar(ctx, cotacaoMoeda)
		if erro != nil {
			t.Fatal(erro)
		}

		var cotacaoMoedaEsperado entidade.CotacaoMoeda

		erro = bancoDados.
			QueryRow(`
				Select 
					id, 
					moeda, 
					de, 
					para, 
					nome,
					valor,
					data 
				from cotacao_moeda
				where id = ?`,
				cotacaoMoeda.Id,
			).
			Scan(
				&cotacaoMoedaEsperado.Id,
				&cotacaoMoedaEsperado.Moeda,
				&cotacaoMoedaEsperado.De,
				&cotacaoMoedaEsperado.Para,
				&cotacaoMoedaEsperado.Nome,
				&cotacaoMoedaEsperado.Valor,
				&cotacaoMoedaEsperado.Data,
			)
		if erro != nil {
			t.Fatal(erro)
		}

		if cotacaoMoeda.Id != cotacaoMoedaEsperado.Id {
			t.Fatalf("resultado: '%v', esperado: '%v'", cotacaoMoeda.Id, cotacaoMoedaEsperado.Id)
		}
	})
}

func obterBancoDadosTeste(t *testing.T) (*sql.DB, error) {
	t.Helper()

	bancoDadosTeste, erro := sql.Open("sqlite3", ":memory:")
	if erro != nil {
		t.Fatal("não foi possível obter uma conexão com o banco de dados")
		return nil, erro
	}

	return bancoDadosTeste, nil
}
