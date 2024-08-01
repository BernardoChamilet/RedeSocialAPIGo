package middlewares

import (
	"api/src/autenticacao"
	"api/src/respostas"
	"log"
	"net/http"
)

//fica no "meio" entre requisição e resposta. Funções que são aplicadas para todas rotas
//logo é útil para uso de tokens
//usando essa abordagem a autenticação das rotas são verificas antes de colocalas no roteador(router). Eu poderia colocar todas e fazer controle de acesso só nos controllers
//esse proximaFunc nessas funções significa que é pra passar para a próxima func aninhada. No final irá executar a função das rotas

// Logger escreve informações da requisição no terminal
func Logger(proximaFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		proximaFunc(w, r)
	}
}

// Autenticar verifica se o usuario apos fazer requisição está autenticado
func Autenticar(proximaFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//vendo se o token é válido
		if erro := autenticacao.ValidarToken(r); erro != nil {
			respostas.Erro(w, http.StatusUnauthorized, erro)
			return
		}
		proximaFunc(w, r)
	}
}
