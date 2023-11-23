package main

import (
	"github.com/mateusgcoelho/poc-golang/config"
	"github.com/mateusgcoelho/poc-golang/internal/database"
	"github.com/mateusgcoelho/poc-golang/internal/router"
)

func main() {
	config.InitDefaultConfig()
	database.InitDatabase()
	
	router.Initialize()

	defer database.GetDbPool().Close()
}