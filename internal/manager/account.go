package manager

import (
	"github.com/Daterdum/OnlineWallet/internal/dao/pg"
	"github.com/Daterdum/OnlineWallet/internal/dto"
	log "github.com/sirupsen/logrus"
)

type AccountManager struct {
	AccountDao pg.AccountDAO //TODO should depend on BaseDAO interface
}

func (manager *AccountManager) Create(dto dto.AccountDTO) (dto.AccountDTO, error) {
	account, err := manager.AccountDao.Create(&dto)
	if err != nil {
		log.Error("Failed to create account")
	}

	return account, err
}

func (manager *AccountManager) Get(userId int) ([]dto.AccountDTO, error) {
	accountDto, err := manager.AccountDao.Get(userId)
	if err != nil {
		log.Warningf("Failed to get account %v", userId)
	}

	return accountDto, err
}

func (manager *AccountManager) Delete(dto dto.AccountDTO) error {
	err := manager.AccountDao.Delete(&dto)
	if err != nil {
		log.Warningf("Failed to delete dto %v", dto)
	}

	return err
}

func (manager *AccountManager) Update(dto dto.AccountDTO) (dto.AccountDTO, error) {
	account, err := manager.AccountDao.Update(&dto)
	if err != nil {
		log.Warningf("Failed to update dto %v", dto)
	}

	return account, err
}
