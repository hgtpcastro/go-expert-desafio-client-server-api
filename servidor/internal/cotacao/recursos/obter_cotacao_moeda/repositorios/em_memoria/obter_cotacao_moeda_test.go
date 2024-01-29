package emmemoria

import (
	"context"
	"testing"

	"github.com/hgtpcastro/go-expert/desafio-client-server-api/servidor/internal/cotacao/recursos/obter_cotacao_moeda/repositorios/erros"
)

func TestObterCotacaoMoedaEmMemoria(t *testing.T) {
	t.Run("obter cotação da moeda [EUR-BRL]", func(t *testing.T) {
		repositorio := NovoRepositorioObterCotacaoMoeda()
		entidade, erro := repositorio.Obter(context.Background(), "EUR-BRL")
		if erro != nil {
			t.Fatal(erro)
		}
		resultado := entidade.Valor
		esperado := "5.3326"

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
		esperado := erros.ErroMoedaNaoInformada

		if resultado != esperado {
			t.Fatalf("resultado: '%T', esperado: '%T'", resultado, esperado)
		}
	})
}
