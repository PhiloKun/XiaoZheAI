package cmd

import (
	"fmt"
	"xiaozhe-ai-backend/config"
	"xiaozhe-ai-backend/db"
	"xiaozhe-ai-backend/router"
)

func Start() {
	config.InitConfig()
	db.DBInitialize()
	router.Initialize()
}

func Clean() {
	fmt.Println("Cleaning...")
}
