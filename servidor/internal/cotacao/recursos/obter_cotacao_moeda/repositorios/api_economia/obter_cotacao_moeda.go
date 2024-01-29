package apieconomia

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/hgtpcastro/go-expert/desafio-client-server-api/servidor/internal/cotacao/entidade"
	"github.com/hgtpcastro/go-expert/desafio-client-server-api/servidor/internal/cotacao/recursos/obter_cotacao_moeda/repositorios/api_economia/mapeadores"
)

const URL_BASE = "https://economia.awesomeapi.com.br/json/last/"

type RepositorioObterCotacaoMoeda struct {
}

func NovoRepositorioObterCotacaoMoeda() *RepositorioObterCotacaoMoeda {
	return &RepositorioObterCotacaoMoeda{}
}

func (r *RepositorioObterCotacaoMoeda) Obter(moeda string) (*entidade.CotacaoMoeda, error) {
	requisicao, erro := http.NewRequest(http.MethodGet, URL_BASE+moeda, nil)
	if erro != nil {
		return nil, erro
	}

	ctx, cancelarRequisicao := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancelarRequisicao()

	requisicao = requisicao.WithContext(ctx)
	resposta, erro := http.DefaultClient.Do(requisicao)
	if erro != nil {
		log.Println(erro)
		return nil, erro
	}
	defer resposta.Body.Close()

	conteudoResposta, erro := io.ReadAll(resposta.Body)
	if erro != nil {
		return nil, erro
	}

	mapeador := mapeadores.NovoMapeadorObterCotacaoMoeda(moeda)
	entidadeMapeada, erro := mapeador.MapearParaEntidade(conteudoResposta)
	if erro != nil {
		return nil, erro
	}

	return entidadeMapeada, nil
}
