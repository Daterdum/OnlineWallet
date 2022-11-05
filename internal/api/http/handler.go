package http

import "net/http"

func handleGet(w http.ResponseWriter, r *http.Request) string {
	return "Successfully handled get"
}
