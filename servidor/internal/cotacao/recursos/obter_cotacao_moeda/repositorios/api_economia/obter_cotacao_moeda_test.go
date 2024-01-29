package apieconomia

import (
	"context"
	"testing"
	"time"

	"github.com/hgtpcastro/go-expert/desafio-client-server-api/servidor/internal/cotacao/recursos/obter_cotacao_moeda/repositorios/erros"
)

func TestObterCotacaoMoedaEmMemoria(t *testing.T) {
	t.Run("obter cotação da moeda [EUR-BRL]", func(t *testing.T) {
		repositorio := NovoRepositorioObterCotacaoMoeda()
		entidade, erro := repositorio.Obter(context.Background(), "EUR-BRL")
		if erro != nil {
			t.Fatal(erro)
		}
		resultado := entidade.De
		esperado := "EUR"

		if resultado != esperado {
			t.Fatalf("resultado: '%s', esperado: '%s'", resultado, esperado)
		}
	})

	t.Run("obter cotação da moeda não encontrada [EXR-BRL]", func(t *testing.T) {
		repositorio := NovoRepositorioObterCotacaoMoeda()
		_, resultado := repositorio.Obter(context.Background(), "EXR-BRL")
		esperado := erros.ErroMoedaNaoEncontrada

		if resultado != esperado {
			t.Fatalf("resultado: '%T', esperado: '%T'", resultado, esperado)
		}
	})

	t.Run("obter cotação da moeda não informada []", func(t *testing.T) {
		repositorio := NovoRepositorioObterCotacaoMoeda()
		_, resultado := repositorio.Obter(context.Background(), "")
		esperado := erros.ErroMoedaNaoEncontrada

		if resultado != esperado {
			t.Fatalf("resultado: '%T', esperado: '%T'", resultado, esperado)
		}
	})

	t.Run("obter cotação da moeda com estouro de tempo", func(t *testing.T) {
		repositorio := NovoRepositorioObterCotacaoMoeda()

		ctx, cancelarRequisicao := context.WithTimeout(context.Background(), 8*time.Millisecond)
		defer cancelarRequisicao()

		_, erro := repositorio.Obter(ctx, "EUR-BRL")
		if erro == nil {
			t.Fatal("era esperado erro de 'context deadline exceeded'")
		}
	})
}
