package erros

const (
	ErroMoedaNaoEncontrada = ErroObterCotacaoMoeda("moeda não encontrada")
	ErroMoedaNaoInformada  = ErroObterCotacaoMoeda("moeda não informada")
)

type ErroObterCotacaoMoeda string

func (e ErroObterCotacaoMoeda) Error() string {
	return string(e)
}
