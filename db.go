package eliab

import (
	"context"
	"database/sql"
	"fmt"

	// Drive mysql...
	_ "github.com/go-sql-driver/mysql"

	"github.com/jgolang/log"
)

// DB connect
var DB *sql.DB

func dbConnect() bool {

	dataSourceName := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", Get.DataBase.User, Get.DataBase.Password, Get.DataBase.Server, Get.DataBase.Port, Get.DataBase.DataBase)

	var err error

	DB, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Error(err)
		return false
	}

	if err = DB.Ping(); err != nil {

		log.Error(err)
		return false
	}

	ctx := context.Background()

	// Ping database to see if it's still alive.
	// Important for handling network issues and long queries.
	err = DB.PingContext(ctx)
	if err != nil {
		log.Fatal("Error pinging database: " + err.Error())
		return false
	}

	err = mysqlVersion()
	if err != nil {
		log.Error(err)
		return false
	}

	return true

}

func mysqlVersion() error {
	query := "select version()"

	row := DB.QueryRow(query)

	var version string

	err := row.Scan(&version)
	if err != nil {
		return err
	}

	log.Info(version)

	return nil
}
