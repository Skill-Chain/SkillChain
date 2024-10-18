package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func connectDB() {

	/*знак "?" добавляет параметры*/
	dataConnectName := "root:@tcp(localhost:5432,/testdb?parseTime=true"

	db, err := sql.Open("postgres", dataConnectName)
	if err != nil {
		err.Error()
	}
	err = db.Ping()
	if err != nil {
		err.Error()
	}
}
