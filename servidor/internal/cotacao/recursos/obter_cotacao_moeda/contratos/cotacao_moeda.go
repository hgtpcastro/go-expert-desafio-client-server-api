package contratos

import "github.com/hgtpcastro/go-expert/desafio-client-server-api/servidor/internal/cotacao/entidade"

type RepositorioCotacaoMoeda interface {
	Obter(moeda string) (*entidade.CotacaoMoeda, error)
}
