package lib

import (
	"fmt"
	"net/http"
)

func Function(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "COol function")
}
