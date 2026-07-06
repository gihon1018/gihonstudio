package main

import (
	"log"

	"gihonstudio/internal/config"
	"gihonstudio/internal/database"
)

func main() {
	appCfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	db, err := database.InitPostgres(&appCfg.Database)
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	if err = sqlDB.Ping(); err != nil {
		log.Fatalf("数据库心跳检测失败: %v", err)
	}

	log.Println("PostgreSQL 连接成功")
}
