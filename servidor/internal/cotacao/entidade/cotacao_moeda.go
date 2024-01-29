package entidade

import (
	"github.com/google/uuid"
)

type CotacaoMoeda struct {
	Id    uuid.UUID
	Moeda string
	De    string
	Para  string
	Nome  string
	Valor string
	Data  string
}

func NovoCotacaoMoeda(moeda, de, para, nome, valor, data string) *CotacaoMoeda {
	return &CotacaoMoeda{
		Id:    uuid.New(),
		Moeda: moeda,
		De:    de,
		Para:  para,
		Nome:  nome,
		Valor: valor,
		Data:  data,
	}
}

func NovoCotacaoMoedaComId(id uuid.UUID, moeda, de, para, nome, valor, data string) *CotacaoMoeda {
	return &CotacaoMoeda{
		Id:    id,
		Moeda: moeda,
		De:    de,
		Para:  para,
		Nome:  nome,
		Valor: valor,
		Data:  data,
	}
}
