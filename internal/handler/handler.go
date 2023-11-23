package handler

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mateusgcoelho/poc-golang/config"
	"github.com/mateusgcoelho/poc-golang/internal/database"
)

var (
	logger *config.Logger
	dbPool *pgxpool.Pool
)

func InitializeHandlers() {
	logger = config.NewLogger("handler");
	dbPool = database.GetDbPool()

	injectUserHandlerDependencies()
}