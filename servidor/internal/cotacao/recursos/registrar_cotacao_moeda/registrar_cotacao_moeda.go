package repositorios

import (
	"context"
	"time"

	"github.com/hgtpcastro/go-expert/desafio-client-server-api/servidor/internal/cotacao/entidade"
	"github.com/hgtpcastro/go-expert/desafio-client-server-api/servidor/internal/cotacao/recursos/registrar_cotacao_moeda/contratos"
)

type RegistarCotacaoMoeda struct {
	repositorio contratos.RepositorioCotacaoMoeda
}

func NovoRegistarCotacaoMoeda(repositorio contratos.RepositorioCotacaoMoeda) *RegistarCotacaoMoeda {
	return &RegistarCotacaoMoeda{
		repositorio: repositorio,
	}
}

func (r *RegistarCotacaoMoeda) Registrar(contexto context.Context, cotacaoMoeda *entidade.CotacaoMoeda) error {
	ctx, cancelarRegistro := context.WithTimeout(contexto, 10*time.Millisecond)
	defer cancelarRegistro()

	if erro := r.repositorio.Registrar(ctx, cotacaoMoeda); erro != nil {
		return erro
	}

	return nil
}
