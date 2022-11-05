package main

import (
	"fmt"
	"github.com/Daterdum/OnlineWallet/internal/api/http/handler"
	"github.com/Daterdum/OnlineWallet/internal/config"
	"github.com/Daterdum/OnlineWallet/internal/dao/pg"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func startServer(addr int) {
	str_addr := strconv.Itoa(addr)
	str_addr = fmt.Sprintf(":%s", str_addr)

	http.HandleFunc("/get", handler.HandleGet)

	err := http.ListenAndServe(str_addr, nil)
	if err != nil {
		log.Errorf("Error happened %e", err)
	}
}

func startDb() {
	pg.Connect()
}

func main() {
	config.LoadConfig()
	log.Infof("Starting service on %s:%v", config.CONFIG.ServiceHost, config.CONFIG.ServicePort)

	startServer(config.CONFIG.ServicePort)
	startDb()
	log.Info("Stopping service")
}
