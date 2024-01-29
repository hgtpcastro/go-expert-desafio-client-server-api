package contratos

import (
	"context"

	"github.com/hgtpcastro/go-expert/desafio-client-server-api/servidor/internal/cotacao/entidade"
)

type RepositorioCotacaoMoeda interface {
	Obter(contexto context.Context, moeda string) (*entidade.CotacaoMoeda, error)
}
