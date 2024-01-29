package main

import (
	"database/sql"
	"log"
	"net/http"

	// sqlite3
	_ "github.com/mattn/go-sqlite3"

	"github.com/hgtpcastro/go-expert/desafio-client-server-api/servidor/internal/cotacao/infra_estrutura/http/obter_cotacao_moeda/manipuladores"
	obter_cotacao_moeda_repositorio "github.com/hgtpcastro/go-expert/desafio-client-server-api/servidor/internal/cotacao/recursos/obter_cotacao_moeda/repositorios/api_economia"
	registrar_cotacao_moeda_repositorio "github.com/hgtpcastro/go-expert/desafio-client-server-api/servidor/internal/cotacao/recursos/registrar_cotacao_moeda/repositorios/banco_de_dados"
)

func main() {
	repositorioObterCotacaoMoeda := obter_cotacao_moeda_repositorio.NovoRepositorioObterCotacaoMoeda()

	bancoDados, erro := sql.Open("sqlite3", "cotacao-moeda.sqlite")
	if erro != nil {
		log.Fatalf("não foi possível acessar o banco de dados, %v", erro)
		return
	}
	defer bancoDados.Close()

	if erro = bancoDados.Ping(); erro != nil {
		log.Fatalf("não foi possível acessar o banco de dados, %v", erro)
	}

	log.Println("banco de dados conectado com sucesso")

	repositorioRegistrarCotacaoMoeda := registrar_cotacao_moeda_repositorio.NovoRepositorioRegistrarCotacaoMoeda(bancoDados)

	manipulador := manipuladores.NovoManipuladorObterCotacaoMoeda(
		repositorioObterCotacaoMoeda,
		repositorioRegistrarCotacaoMoeda,
	)

	if erro := http.ListenAndServe(":3000", manipulador); erro != nil {
		log.Fatalf("não foi possível escutar na porta 3000, %v", erro)
	}
}
