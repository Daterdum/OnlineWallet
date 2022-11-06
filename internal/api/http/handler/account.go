package handler

import (
	"encoding/json"
	"github.com/Daterdum/OnlineWallet/internal/dto"
	"github.com/Daterdum/OnlineWallet/internal/manager"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type AccountHandler struct {
	AccountManager manager.AccountManager
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
	accounts, err := handler.AccountManager.Get(dto.UserId)
	if err != nil {
		log.Warningf("Failed to get accounts %v", err)
	}
	log.Debug(accounts)
}

func (handler *AccountHandler) create(w http.ResponseWriter, dto dto.AccountDTO) {
	resp, err := handler.AccountManager.Create(dto)
	if err != nil {
		log.Warningf("Failed to create dto %v", err)
	}
	w.Write([]byte(resp.Currency))
	log.Debug(resp)
}

func (handler *AccountHandler) update(w http.ResponseWriter, dto dto.AccountDTO) {
	resp, err := handler.AccountManager.Update(dto)
	if err != nil {
		log.Warningf("Failed to update dto %v", err)
	}
	w.Write([]byte(resp.Currency))

	log.Debug(resp)
}

func (handler *AccountHandler) delete(w http.ResponseWriter, dto dto.AccountDTO) {
	resp := handler.AccountManager.Delete(dto)
	log.Debug(resp)
}
