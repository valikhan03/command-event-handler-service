package database

import(
	"fmt"
	"log"
	"github.com/jmoiron/sqlx"

	"command-event-handler-service/models"
)

var DBConn *sqlx.DB

func InitDatabase() {
	var err error
	connStr := fmt.Sprintf("host=%s, port=%d, user=%s, name=%s, password=%s, sslmode=%s", 
					models.DBConfs.Host, models.DBConfs.Port, models.DBConfs.User, models.DBConfs.DBName, models.DBConfs.Password, models.DBConfs.SSLMode)

	DBConn, err = sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

}