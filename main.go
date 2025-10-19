package main

import (
	"fmt"
	"gin-ayo/config"
	"gin-ayo/database/connectors"
	storage "gin-ayo/pkg/supabase"
	"gin-ayo/pkg/validator"
	"gin-ayo/routes"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
)

var pgsqldb *gorm.DB

func init() {
	var err error

	_ = godotenv.Load()

	err = config.CheckEnv()
	if err != nil {
		fmt.Println(err)
	}

	pgsqlConn := connectors.PgSQLConn{
		DbHost:     config.GetEnv("DB_PGSQL_HOST", ""),
		DbPort:     config.GetEnv("DB_PGSQL_PORT", ""),
		DbDatabase: config.GetEnv("DB_PGSQL_DATABASE", ""),
		DbUsername: config.GetEnv("DB_PGSQL_USERNAME", ""),
		DbPassword: config.GetEnv("DB_PGSQL_PASSWORD", ""),
	}

	dbConn, err := connectors.NewPgSQLConn(pgsqlConn)
	if err != nil {
		fmt.Println(err)
	}

	if dbConn == nil {
		fmt.Println("failed connect to pgsqldb database")
	}

	db, err := dbConn.DB()
	if err != nil {
		fmt.Println("failed connect to database")
	}
	db.SetMaxOpenConns(500)
	db.SetMaxIdleConns(100)

	pgsqldb = dbConn
}

func main() {
	port := config.GetEnv("PORT", "8080")
	storage.InitSupabaseS3()
	validator.RegisterCustomValidations()
	route := routes.NewRoute(pgsqldb)
	app := route.SetupRoutes()

	err := app.Run(":" + port)
	log.Printf("connect to http://localhost:%s/", port)
	if err != nil {
		log.Println("Failed To Start System")
		panic("Failed To Start")
	}
}
