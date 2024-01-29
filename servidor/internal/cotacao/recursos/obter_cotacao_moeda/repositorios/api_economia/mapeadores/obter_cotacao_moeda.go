package mapeadores

import (
	"encoding/json"

	"github.com/hgtpcastro/go-expert/desafio-client-server-api/servidor/internal/cotacao/entidade"
)

type MapeadorObterCotacaoMoeda struct {
	moeda string
}

func NovoMapeadorObterCotacaoMoeda(moeda string) *MapeadorObterCotacaoMoeda {
	return &MapeadorObterCotacaoMoeda{
		moeda: moeda,
	}
}

func (m *MapeadorObterCotacaoMoeda) MapearParaEntidade(de []byte) (*entidade.CotacaoMoeda, error) {
	var dadoGenerico map[string]any

	if erro := json.Unmarshal(de, &dadoGenerico); erro != nil {
		return nil, erro
	}

	chaveMoeda := m.moeda[:3] + m.moeda[4:]
	cotacaoMoeda := dadoGenerico[chaveMoeda].(map[string]any)

	// for key, value := range birds {
	// 	// Each value is an `any` type, that is type asserted as a string
	// 	fmt.Println(key, value.(string))
	// }

	para := entidade.NovoCotacaoMoeda(
		m.moeda,
		cotacaoMoeda["code"].(string),
		cotacaoMoeda["codein"].(string),
		cotacaoMoeda["name"].(string),
		cotacaoMoeda["bid"].(string),
		cotacaoMoeda["create_date"].(string),
	)

	return para, nil
}
