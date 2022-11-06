package pg

import (
	"context"
	"github.com/Daterdum/OnlineWallet/internal/dto"
	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"
)

type AccountDAO struct {
	DbPool *pgxpool.Pool
}

func (dao *AccountDAO) Create(account *dto.AccountDTO) (dto.AccountDTO, error) {
	err := dao.DbPool.QueryRow(
		context.Background(),
		"INSERT INTO wallet.account (user_id, amount, currency) VALUES ($1, $2, $3) RETURNING id;",
		account.UserId, account.Amount, account.Currency).Scan(&account.Id)

	if err != nil {
		log.Errorf("QueryRow failed: %v\n", err)
	}

	return *account, err
}

func (dao *AccountDAO) Get(userId int) ([]dto.AccountDTO, error) {
	rows, err := dao.DbPool.Query(
		context.Background(),
		"SELECT id, user_id, amount, currency FROM wallet.account WHERE account.user_id = $1",
		userId,
	)
	if err != nil {
		log.Errorf("QueryRow failed: %v\n", err)
	}

	var result []dto.AccountDTO
	rows.Scan(&result)

	return result, err
}

func (dao *AccountDAO) Delete(account *dto.AccountDTO) error {
	_, err := dao.DbPool.Exec(
		context.Background(),
		"DELETE FROM wallet.account WHERE account.id = $1;",
		account.Id,
	)

	if err != nil {
		log.Errorf("Deletion failed: %v\n", err)
	}

	return err
}

func (dao *AccountDAO) Update(account *dto.AccountDTO) (dto.AccountDTO, error) {
	err := dao.DbPool.QueryRow(
		context.Background(),
		"UPDATE wallet.account (amount) SET VALUES ($1) WHERE account.id = $3 RETURNING id, amount, currency, user_id;",
		account.Amount).Scan(&account.Id, &account.Amount, &account.Currency, &account.UserId)

	if err != nil {
		log.Errorf("QueryRow failed: %v\n", err)
	}

	return *account, err
}
