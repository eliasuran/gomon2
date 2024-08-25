package main

import (
	"net/http"

	"github.com/eliasuran/gomon2/test/lib"
)

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
    lib.Function(w, r)
  })
  http.ListenAndServe(":8080", mux)
}
