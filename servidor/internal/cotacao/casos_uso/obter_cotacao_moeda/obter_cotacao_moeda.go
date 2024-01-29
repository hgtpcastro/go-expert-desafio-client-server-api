package obtercotacaomoeda

import (
	"context"

	"github.com/google/uuid"
	"github.com/hgtpcastro/go-expert/desafio-client-server-api/servidor/internal/cotacao/dtos"
	"github.com/hgtpcastro/go-expert/desafio-client-server-api/servidor/internal/cotacao/entidade"
	obter_cotacao_moeda "github.com/hgtpcastro/go-expert/desafio-client-server-api/servidor/internal/cotacao/recursos/obter_cotacao_moeda"
	obter_cotacao_moeda_contratos "github.com/hgtpcastro/go-expert/desafio-client-server-api/servidor/internal/cotacao/recursos/obter_cotacao_moeda/contratos"
	registrar_cotacao_moeda "github.com/hgtpcastro/go-expert/desafio-client-server-api/servidor/internal/cotacao/recursos/registrar_cotacao_moeda"
	registrar_cotacao_moeda_contratos "github.com/hgtpcastro/go-expert/desafio-client-server-api/servidor/internal/cotacao/recursos/registrar_cotacao_moeda/contratos"
)

type CasoUsoObterCotacaoMoeda struct {
	repositorioObterCotacaoMoeda     obter_cotacao_moeda_contratos.RepositorioCotacaoMoeda
	repositorioRegistrarCotacaoMoeda registrar_cotacao_moeda_contratos.RepositorioCotacaoMoeda

	obterCotacaoMoeda     *obter_cotacao_moeda.ObterCotacaoMoeda
	registrarCotacaoMoeda *registrar_cotacao_moeda.RegistarCotacaoMoeda
}

func NovoCasoUsoObterCotacaoMoeda(
	repositorioObterCotacaoMoeda obter_cotacao_moeda_contratos.RepositorioCotacaoMoeda,
	repositorioRegistrarCotacaoMoeda registrar_cotacao_moeda_contratos.RepositorioCotacaoMoeda) *CasoUsoObterCotacaoMoeda {
	c := &CasoUsoObterCotacaoMoeda{
		repositorioObterCotacaoMoeda:     repositorioObterCotacaoMoeda,
		repositorioRegistrarCotacaoMoeda: repositorioRegistrarCotacaoMoeda,
	}
	c.ConfigurarDependencias()
	return c
}

func (s *CasoUsoObterCotacaoMoeda) ConfigurarDependencias() {
	s.obterCotacaoMoeda = obter_cotacao_moeda.NovoObterCotacaoMoeda(s.repositorioObterCotacaoMoeda)
	if s.repositorioRegistrarCotacaoMoeda != nil {
		s.registrarCotacaoMoeda = registrar_cotacao_moeda.NovoRegistarCotacaoMoeda(s.repositorioRegistrarCotacaoMoeda)
	}
}

func (s *CasoUsoObterCotacaoMoeda) Obter(contexto context.Context, moeda string) (dtos.ObterCotacaoMoedaDto, error) {
	obterCotacaoMoedaDto, erro := s.obterCotacaoMoeda.Obter(contexto, moeda)
	if erro != nil {
		return dtos.ObterCotacaoMoedaDto{}, erro
	}

	if s.registrarCotacaoMoeda == nil {
		return obterCotacaoMoedaDto, nil
	}

	id, erro := uuid.Parse(obterCotacaoMoedaDto.Id)
	if erro != nil {
		return dtos.ObterCotacaoMoedaDto{}, erro
	}

	cotacoMoeda := entidade.NovoCotacaoMoedaComId(
		id,
		obterCotacaoMoedaDto.Moeda,
		obterCotacaoMoedaDto.De,
		obterCotacaoMoedaDto.Para,
		obterCotacaoMoedaDto.Nome,
		obterCotacaoMoedaDto.Valor,
		obterCotacaoMoedaDto.Data,
	)

	if erro = s.registrarCotacaoMoeda.Registrar(contexto, cotacoMoeda); erro != nil {
		return dtos.ObterCotacaoMoedaDto{}, erro
	}

	return obterCotacaoMoedaDto, nil
}
