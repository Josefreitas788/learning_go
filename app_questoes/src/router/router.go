package router

import (
	"api/src/router/rotas"

	"github.com/gorilla/mux"
)

// GERAR - Gera as rotas da aplicação
func Gerar() *mux.Router {

	r := mux.NewRouter()
	return rotas.Configutar(r)

}
