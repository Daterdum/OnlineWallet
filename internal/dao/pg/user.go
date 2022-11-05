package pg

import (
	"context"
	"github.com/Daterdum/OnlineWallet/internal/dto"
	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"
)

type UserDAO struct {
	dbPool pgxpool.Pool
}

func (dao *UserDAO) Create(user *dto.UserDTO) error {
	err := dao.dbPool.QueryRow(context.Background(), "INSERT INTO wallet.user (name) VALUES ('$1') RETURNING id", user.Name).Scan(&user.Id)
	if err != nil {
		log.Errorf("QueryRow failed: %v\n", err)
	}
	return err
}
