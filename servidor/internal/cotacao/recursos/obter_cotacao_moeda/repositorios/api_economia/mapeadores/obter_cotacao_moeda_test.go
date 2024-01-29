package mapeadores

import (
	"testing"

	"github.com/hgtpcastro/go-expert/desafio-client-server-api/servidor/internal/cotacao/entidade"
)

func TestMapeadorObterCotacaoMoeda(t *testing.T) {
	t.Run("mapear resposta da requisicao para entidade", func(t *testing.T) {
		//cotacaoDto := &apieconomia.ObterCotacaoMoedaResposta{}

		conteudoResposta := []byte(`
		{
			"USDBRL": {
				"code": "USD",
				"codein": "BRL",
				"name": "Dólar Americano/Real Brasileiro",
				"high": "4.92",
				"low": "4.9022",
				"varBid": "0.0008",
				"pctChange": "0.02",
				"bid": "4.9168",
				"ask": "4.9178",
				"timestamp": "1706304601",
				"create_date": "2024-01-26 18:30:01"
			}
		}
		`)

		var entidade *entidade.CotacaoMoeda

		mapeador := NovoMapeadorObterCotacaoMoeda("USD-BRL")
		entidade, erro := mapeador.MapearParaEntidade(conteudoResposta)
		if erro != nil {
			t.Fatal(erro)
		}

		if entidade == nil {
			t.Fatal("esperava-se que a entidade estivesse sido mapeada")
		}

		if entidade.Valor != "4.9168" {
			t.Fatal("esperava-se que a entidade estivesse sido mapeada com os valores corretamente")
		}

	})
}
