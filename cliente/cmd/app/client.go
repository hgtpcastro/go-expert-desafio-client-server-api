package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type CotacaoMoeda struct {
	Nome  string `jason:"nome"`
	Valor string `json:"valor"`
}

const URL_BASE = "http://localhost:3000/obter-cotacao-moeda/USD-BRL"

func main() {
	ctx, cancelarRequisicao := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancelarRequisicao()

	requisicao, erro := http.NewRequestWithContext(ctx, http.MethodGet, URL_BASE, nil)
	if erro != nil {
		panic(erro)
	}

	resposta, erro := http.DefaultClient.Do(requisicao)
	if erro != nil {
		log.Println(erro)
		panic(erro)
	}
	defer resposta.Body.Close()

	conteudo, erro := io.ReadAll(resposta.Body)
	if erro != nil {
		panic(erro)
	}

	var cotacaoMoeda CotacaoMoeda
	erro = json.Unmarshal(conteudo, &cotacaoMoeda)
	if erro != nil {
		fmt.Fprintf(os.Stderr, "erro ao fazer parse da resposta: %v\n", erro)
		panic(erro)
	}

	file, erro := os.Create("cotacao.txt")
	if erro != nil {
		fmt.Fprintf(os.Stderr, "erro ao criar arquivo: %v\n", erro)
		panic(erro)
	}
	defer file.Close()

	_, erro = file.WriteString(fmt.Sprintf("%s: %s\n", cotacaoMoeda.Nome, cotacaoMoeda.Valor))
	if erro != nil {
		fmt.Fprintf(os.Stderr, "Erro ao escrever no arquivo: %v\n", erro)
		panic(erro)
	}

	fmt.Println("Arquivo criado com sucesso!")
	fmt.Printf("%s: %s\n", cotacaoMoeda.Nome, cotacaoMoeda.Valor)
}
