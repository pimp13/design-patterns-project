package main

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/pimp13/go-react-project/cmd/api"
	"github.com/pimp13/go-react-project/config"
	"github.com/pimp13/go-react-project/db"
	"log"
)

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatalf("error to ping db: %s", err)
	}

	log.Println("Database successfully connected!")
}
func main() {
	// Create instance database
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}
	initStorage(db)

	// Create instance API server
	server := api.NewAPIServer(config.Envs.Port, db)
	// Run Application
	if err := server.Run(); err != nil {
		log.Fatalf("Error running server: %v", err)
	}
}
