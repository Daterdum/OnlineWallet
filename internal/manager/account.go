package manager

import (
	"github.com/Daterdum/OnlineWallet/internal/dao/pg"
	"github.com/Daterdum/OnlineWallet/internal/dto"
)

type AccountManager struct {
	accountDao pg.AccountDAO
}

func (manager *AccountManager) Create(dto dto.AccountDTO) dto.AccountDTO {
	resp := manager.accountDao.Create(dto)
	return resp
}

func (manager *AccountManager) Get(userId int) []dto.AccountDTO {
	resp := manager.accountDao.Get(userId)
	return resp
}

func (manager *AccountManager) Delete(dto dto.AccountDTO) {

}

func (manager *AccountManager) Update(dto dto.AccountDTO) {

}
