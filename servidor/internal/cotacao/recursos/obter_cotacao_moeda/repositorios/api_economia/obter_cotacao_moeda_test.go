package apieconomia

import (
	"testing"
)

func TestObterCotacaoMoedaEmMemoria(t *testing.T) {
	t.Run("obter cotação da moeda [EUR-BRL]", func(t *testing.T) {
		repositorio := NovoRepositorioObterCotacaoMoeda()
		entidade, erro := repositorio.Obter("EUR-BRL")
		if erro != nil {
			t.Fatal(erro)
		}
		resultado := entidade.De
		esperado := "EUR"

		if resultado != esperado {
			t.Fatalf("resultado: '%s', esperado: '%s'", resultado, esperado)
		}
	})

	// t.Run("obter cotação da moeda não encontrada [EXR-BRL]", func(t *testing.T) {
	// 	repositorio := NovoRepositorioObterCotacaoMoeda()
	// 	_, resultado := repositorio.Obter("EXR-BRL")
	// 	esperado := erros.ErroMoedaNaoEncontrada

	// 	if resultado != esperado {
	// 		t.Fatalf("resultado: '%T', esperado: '%T'", resultado, esperado)
	// 	}
	// })

	// t.Run("obter cotação da moeda não informada []", func(t *testing.T) {
	// 	repositorio := NovoRepositorioObterCotacaoMoeda()
	// 	_, resultado := repositorio.Obter("")
	// 	esperado := erros.ErroMoedaNaoEncontrada

	// 	if resultado != esperado {
	// 		t.Fatalf("resultado: '%T', esperado: '%T'", resultado, esperado)
	// 	}
	// })
}
