package cmd

import (
	"fmt"
	"github.com/philokun/xiaozhe-ai-backend/config"
	"github.com/philokun/xiaozhe-ai-backend/db"
	"github.com/philokun/xiaozhe-ai-backend/router"
)

func Start() {
	config.InitConfig()
	db.DBInitialize()
	router.Initialize()
}

func Clean() {
	fmt.Println("Cleaning...")
}
