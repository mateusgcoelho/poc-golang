package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mateusgcoelho/poc-golang/config"
)

var (
	dbPool *pgxpool.Pool
)

func InitDatabase() {
	db, err := connectToDatabase() 
	logger := config.NewLogger("database-config")

	if err != nil {
		logger.Errorf("unable to connect at database: %v", err)
		os.Exit(1);
		return
	}

	dbPool = db
}

func connectToDatabase() (*pgxpool.Pool, error) {
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	
    if err != nil {
        return nil, fmt.Errorf("unable to connect to database: %v", err)
    }

    if err := dbpool.Ping(context.Background()); err != nil {
        dbpool.Close()
        return nil, fmt.Errorf("error pinging database")
    }

    return dbpool, nil
}

func GetDbPool() *pgxpool.Pool {
	return dbPool
}
