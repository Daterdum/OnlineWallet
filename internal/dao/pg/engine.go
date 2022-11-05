package pg

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	log "github.com/sirupsen/logrus"
	"os"
)

func Connect() {
	log.Infoln("Started accessing db")
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var id int
	err = conn.QueryRow(context.Background(), "SELECT id from message.message LIMIT $1", 1).Scan(&id)
	if err != nil {
		log.Errorf("QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	log.Infoln(id)
}
