package handler

import (
	"fmt"
	"github.com/Daterdum/OnlineWallet/internal/dao/pg"
	"net/http"
)

func HandleGet(w http.ResponseWriter, r *http.Request) {
	pg.Connect()
	fmt.Fprintln(w, "Successfully handled get")
}
