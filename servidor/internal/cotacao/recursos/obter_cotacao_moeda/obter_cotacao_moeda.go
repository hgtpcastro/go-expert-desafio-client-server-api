package obtercotacaomoeda

import (
	"context"
	"time"

	"github.com/hgtpcastro/go-expert/desafio-client-server-api/servidor/internal/cotacao/dtos"
	"github.com/hgtpcastro/go-expert/desafio-client-server-api/servidor/internal/cotacao/recursos/obter_cotacao_moeda/contratos"
)

type ObterCotacaoMoeda struct {
	repositorio contratos.RepositorioCotacaoMoeda
}

func NovoObterCotacaoMoeda(repositorioCotacao contratos.RepositorioCotacaoMoeda) *ObterCotacaoMoeda {
	return &ObterCotacaoMoeda{
		repositorio: repositorioCotacao,
	}
}

func (o *ObterCotacaoMoeda) Obter(contexto context.Context, moeda string) (dtos.ObterCotacaoMoedaDto, error) {
	ctx, cancelarRequisicao := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancelarRequisicao()

	entidade, erro := o.repositorio.Obter(ctx, moeda)
	if erro != nil {
		return dtos.ObterCotacaoMoedaDto{}, erro
	}

	cotacaoDto := dtos.ObterCotacaoMoedaDto{
		Id:    entidade.Id.String(),
		Moeda: entidade.Moeda,
		De:    entidade.De,
		Para:  entidade.Para,
		Nome:  entidade.Nome,
		Valor: entidade.Valor,
		Data:  entidade.Data,
	}
	return cotacaoDto, nil
}
