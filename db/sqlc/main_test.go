package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries

const (
	dbDriver = "postgres"
	dbSoure  = "postgresql://root:secret@156.253.5.166:5432/simple_bank?sslmode=disable"
)

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSoure)
	if err != nil {
		log.Fatal("cannot connet to db: ", err.Error())
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
