package pg

import (
	"github.com/Daterdum/OnlineWallet/internal/dto"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AccountDAO struct {
	dbPool pgxpool.Pool
}

func (dao *AccountDAO) Create(account dto.AccountDTO) dto.AccountDTO {

}

func (dao *AccountDAO) Get(userId int) []dto.AccountDTO {

}

func (dao *AccountDAO) Delete(account dto.AccountDTO) {

}

func (dao *AccountDAO) Update(account dto.AccountDTO) dto.AccountDTO {

}
