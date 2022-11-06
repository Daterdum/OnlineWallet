package main

import (
	"fmt"
	"github.com/Daterdum/OnlineWallet/internal/api/http/handler"
	"github.com/Daterdum/OnlineWallet/internal/config"
	"github.com/Daterdum/OnlineWallet/internal/dao/pg"
	"github.com/Daterdum/OnlineWallet/internal/manager"
	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func registerManagers(DbPool *pgxpool.Pool) manager.AccountManager {
	return manager.AccountManager{AccountDao: pg.AccountDAO{DbPool: DbPool}}
}

func registerHandlers(dbPool *pgxpool.Pool) {
	accountManager := registerManagers(dbPool)
	accountHandler := handler.AccountHandler{accountManager}
	http.HandleFunc("/account", accountHandler.Router)
}

func startServer(addr int, dbPool *pgxpool.Pool) {
	str_addr := strconv.Itoa(addr)
	str_addr = fmt.Sprintf(":%s", str_addr)

	registerHandlers(dbPool)

	err := http.ListenAndServe(str_addr, nil)
	if err != nil {
		log.Errorf("Error happened %e", err)
	}
}

func startDb() *pgxpool.Pool {
	return pg.SetupPool()
}

func main() {
	log.SetLevel(log.DebugLevel)
	config.LoadConfig()
	log.Infof("Starting service on %s:%v", config.CONFIG.ServiceHost, config.CONFIG.ServicePort)
	defer log.Info("Stopping service")
	dbPool := startDb()
	startServer(
		config.CONFIG.ServicePort,
		dbPool,
	)
	defer dbPool.Close()
}
