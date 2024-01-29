package emmemoria

import (
	"context"
	"strings"

	"github.com/hgtpcastro/go-expert/desafio-client-server-api/servidor/internal/cotacao/entidade"
	"github.com/hgtpcastro/go-expert/desafio-client-server-api/servidor/internal/cotacao/recursos/obter_cotacao_moeda/repositorios/erros"
)

type RepositorioObterCotacaoMoeda struct {
	listaCotacoes map[string]*entidade.CotacaoMoeda
}

func NovoRepositorioObterCotacaoMoeda() *RepositorioObterCotacaoMoeda {
	r := &RepositorioObterCotacaoMoeda{}
	r.carregarCotacoes()
	return r
}

func (r *RepositorioObterCotacaoMoeda) Obter(contexto context.Context, moeda string) (*entidade.CotacaoMoeda, error) {
	if strings.TrimSpace(moeda) == "" {
		return nil, erros.ErroMoedaNaoInformada
	}

	entidade, existe := r.listaCotacoes[moeda]
	if !existe {
		return nil, erros.ErroMoedaNaoEncontrada
	}

	return entidade, nil
}

func (r *RepositorioObterCotacaoMoeda) carregarCotacoes() {
	r.listaCotacoes = map[string]*entidade.CotacaoMoeda{
		"USD-BRL": entidade.NovoCotacaoMoeda(
			"USD-BRL",
			"USD",
			"BRL",
			"Dólar Americano/Real Brasileiro",
			"4.9165",
			"2024-01-26 18:59:33",
		),
		"EUR-BRL": entidade.NovoCotacaoMoeda(
			"EUR-BRL",
			"EUR",
			"BRL",
			"Euro/Real Brasileiro",
			"5.3326",
			"2024-01-26 18:59:33",
		),
	}
}
