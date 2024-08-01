package respostas

//pacote responsavel por padronizar as resposta que a api dará ao frontend
import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON recebe um status e um conteúdo e retorna um Json ao cliente
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if dados != nil {
		//transformando em json
		if erro := json.NewEncoder(w).Encode(dados); erro != nil {
			log.Fatal(erro)
		}
	}
}

// Erro retorna um json de erro ao cliente
func Erro(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(), //retorna erro do go
	})
}
