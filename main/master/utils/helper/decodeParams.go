package helper

import (
	"net/http"

	"github.com/gorilla/mux"
)

// DecodeParams -> decode paramater from url
func DecodeParams(val string, r *http.Request) string {
	param := mux.Vars(r)
	return param[val]
}
