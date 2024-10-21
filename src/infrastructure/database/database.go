package database

import (
	"database/sql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var DbClient *sql.DB

func ConnectDB() {
	/*знак "?" добавляет параметры*/
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	DataConnectName, existsDn := os.LookupEnv("DATA_COONECT_NAME")
	NameDataBase, existsDB := os.LookupEnv("NAME_DATA_BASE")
	if !existsDn || !existsDB {
		log.Fatal("Error: invalid parameters in dir .env")
	}

	db, err := sql.Open(NameDataBase, DataConnectName)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	DbClient = db
}
