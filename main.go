package main

import (
	"net/http"

	_ "github.com/lib/pq"
	"std/github.com/jpodlasnisky/aluragolang/fundamentosweb/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
