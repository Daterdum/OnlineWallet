package handler

import (
	"fmt"
	"net/http"
)

func HandleGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Successfully handled get")
}
