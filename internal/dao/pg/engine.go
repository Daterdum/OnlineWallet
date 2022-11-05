package pg

import (
	"context"
	"fmt"
	"github.com/Daterdum/OnlineWallet/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"
	"os"
)

func SetupPool() *pgxpool.Pool {
	dbPool, err := pgxpool.New(context.Background(), config.CONFIG.DatabaseUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return dbPool
}

func Connect(dbPool pgxpool.Pool) int {
	var id int
	err := dbPool.QueryRow(context.Background(), "SELECT id from message.message LIMIT $1", 1).Scan(&id)
	if err != nil {
		log.Errorf("QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	log.Infoln(id)

	return id
}
