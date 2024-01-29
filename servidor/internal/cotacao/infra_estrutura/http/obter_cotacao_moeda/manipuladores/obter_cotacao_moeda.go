package manipuladores

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	obter_cotacao_moeda_caso_uso "github.com/hgtpcastro/go-expert/desafio-client-server-api/servidor/internal/cotacao/casos_uso/obter_cotacao_moeda"
	obter_cotacao_moeda_contratos "github.com/hgtpcastro/go-expert/desafio-client-server-api/servidor/internal/cotacao/recursos/obter_cotacao_moeda/contratos"
	registrar_cotacao_moeda_contratos "github.com/hgtpcastro/go-expert/desafio-client-server-api/servidor/internal/cotacao/recursos/registrar_cotacao_moeda/contratos"
)

const URL = "/obter-cotacao-moeda/"

type ManipuladorObterCotacaoMoeda struct {
	repositorioObterCotacaoMoeda     obter_cotacao_moeda_contratos.RepositorioCotacaoMoeda
	repositorioRegistrarCotacaoMoeda registrar_cotacao_moeda_contratos.RepositorioCotacaoMoeda
}

func NovoManipuladorObterCotacaoMoeda(
	repositorioObterCotacaoMoeda obter_cotacao_moeda_contratos.RepositorioCotacaoMoeda,
	repositorioRegistrarCotacaoMoeda registrar_cotacao_moeda_contratos.RepositorioCotacaoMoeda,
) *ManipuladorObterCotacaoMoeda {
	return &ManipuladorObterCotacaoMoeda{
		repositorioObterCotacaoMoeda:     repositorioObterCotacaoMoeda,
		repositorioRegistrarCotacaoMoeda: repositorioRegistrarCotacaoMoeda,
	}
}

func (m *ManipuladorObterCotacaoMoeda) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	moeda := r.URL.Path[len(URL):]

	casoUsoObterCotacaoMoeda := obter_cotacao_moeda_caso_uso.NovoCasoUsoObterCotacaoMoeda(
		m.repositorioObterCotacaoMoeda,
		m.repositorioRegistrarCotacaoMoeda,
	)

	cotacaoDto, erro := casoUsoObterCotacaoMoeda.Obter(context.Background(), moeda)

	if erro != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, erro.Error())
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cotacaoDto)
}
