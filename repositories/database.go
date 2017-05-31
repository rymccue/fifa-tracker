package repositories

import (
	"database/sql"

	"github.com/knq/dburl"
	_ "github.com/lib/pq"
)

var db *sql.DB

func CreateDatabase(connectionString string) error {
	var err error
	db, err = dburl.Open(connectionString)
	return err
}
