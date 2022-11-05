package pg

import (
	"context"
	"github.com/Daterdum/OnlineWallet/internal/dto"
	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"
	"os"
)

type AccountDAO struct {
	dbPool pgxpool.Pool
}

func (dao *AccountDAO) Create(account dto.AccountDTO) error {
	err := dao.dbPool.QueryRow(
		context.Background(),
		"INSERT INTO wallet.account (user_id, amount, currency) VALUES (%1, %2, %3); RETURNING id",
		account.UserId, account.Amount, account.Currency).Scan(&account.Id)

	if err != nil {
		log.Errorf("QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	return err
}

func (dao *AccountDAO) Get(userId int) {

}

func (dao *AccountDAO) Delete(account dto.AccountDTO) {

}

func (dao *AccountDAO) Update(account dto.AccountDTO) {

}
