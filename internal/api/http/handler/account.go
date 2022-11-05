package handler

import (
	"encoding/json"
	"github.com/Daterdum/OnlineWallet/internal/dto"
	"github.com/Daterdum/OnlineWallet/internal/manager"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type AccountHandler struct {
	accountManager manager.AccountManager
}

type AccountRequest struct {
	Method string         `json:"method"`
	Data   dto.AccountDTO `json:"data"`
}

func (handler *AccountHandler) Router(w http.ResponseWriter, r *http.Request) {
	if r.Body != nil {
		defer r.Body.Close()
	}
	acc := AccountRequest{}
	err := json.NewDecoder(r.Body).Decode(&acc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	switch acc.Method {
	case "create":
		handler.create(w, acc.Data)
	case "get":
		handler.get(w, acc.Data)
	}
}

func (handler *AccountHandler) get(w http.ResponseWriter, dto dto.AccountDTO) {
	resp := handler.accountManager.Get(dto.UserId)
	log.Debug(resp)
}

func (handler *AccountHandler) create(w http.ResponseWriter, dto dto.AccountDTO) {
	resp := handler.accountManager.Create(dto)
	log.Debug(resp)
}
