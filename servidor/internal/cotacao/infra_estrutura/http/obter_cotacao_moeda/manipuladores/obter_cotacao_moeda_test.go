package manipuladores

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hgtpcastro/go-expert/desafio-client-server-api/servidor/internal/cotacao/dtos"
	emmemoria "github.com/hgtpcastro/go-expert/desafio-client-server-api/servidor/internal/cotacao/recursos/obter_cotacao_moeda/repositorios/em_memoria"
)

func TestObterCotacao(t *testing.T) {
	t.Run("obter cotação da moeda [USD-BRL]", func(t *testing.T) {
		requisicao := novaRequisicaoObterCotacaoMoeda("USD-BRL")
		resposta := httptest.NewRecorder()

		repositorio := emmemoria.NovoRepositorioObterCotacaoMoeda()
		manipulador := NovoManipuladorObterCotacaoMoeda(repositorio, nil)
		manipulador.ServeHTTP(resposta, requisicao)

		verificarRespostaRequisicao(t, resposta.Body, "4.9165")
	})

	t.Run("obter cotação da moeda [EUR-BRL]", func(t *testing.T) {
		requisicao := novaRequisicaoObterCotacaoMoeda("EUR-BRL")
		resposta := httptest.NewRecorder()

		repositorio := emmemoria.NovoRepositorioObterCotacaoMoeda()
		manipulador := NovoManipuladorObterCotacaoMoeda(repositorio, nil)
		manipulador.ServeHTTP(resposta, requisicao)

		verificarRespostaRequisicao(t, resposta.Body, "5.3326")
	})

}

func novaRequisicaoObterCotacaoMoeda(moeda string) *http.Request {
	requisicao, _ := http.NewRequest(http.MethodGet, URL+moeda, nil)
	return requisicao
}

func verificarRespostaRequisicao(t *testing.T, resultado *bytes.Buffer, esperado string) {
	t.Helper()

	conteudo, erro := io.ReadAll(resultado)
	if erro != nil {
		t.Fatal(erro)
	}

	var cotacaoDto dtos.ObterCotacaoMoedaDto
	erro = json.Unmarshal(conteudo, &cotacaoDto)
	if erro != nil {
		t.Fatal(erro)
	}

	if cotacaoDto.Valor != esperado {
		t.Fatalf("resultado: '%v', esperado: '%v'", resultado, esperado)
	}
}
