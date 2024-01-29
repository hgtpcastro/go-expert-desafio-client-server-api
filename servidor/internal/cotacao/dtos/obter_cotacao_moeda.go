package dtos

type ObterCotacaoMoedaDto struct {
	Id    string `json:"id"`
	Moeda string `json:"moeda"`
	De    string `json:"de"`
	Para  string `json:"para"`
	Nome  string `json:"nome"`
	Valor string `json:"valor"`
	Data  string `json:"data"`
}
