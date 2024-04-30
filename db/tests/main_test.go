package db_test

import (
	"database/sql"
	"log"
	"os"
	"testing"

	db "github.com/adityanuriskandar17/fingreat-backend/db/sqlc"
	_ "github.com/lib/pq"
)

var testQuery *db.Queries

// db connection
func TestMain(m *testing.M) {
	conn, err := sql.Open("postgres", "postgres://root:secret@localhost:5454/fingreat_db?sslmode=disable")
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQuery = db.New(conn)

	os.Exit(m.Run())
}
