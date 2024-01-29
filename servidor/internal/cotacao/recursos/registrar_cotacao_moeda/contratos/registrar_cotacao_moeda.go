package contratos

import (
	"context"

	"github.com/hgtpcastro/go-expert/desafio-client-server-api/servidor/internal/cotacao/entidade"
)

type RepositorioCotacaoMoeda interface {
	Registrar(contexto context.Context, cotacaoMoeda *entidade.CotacaoMoeda) error
}
