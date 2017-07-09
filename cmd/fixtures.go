package main

import (
	"database/sql"
	"log"
	"gopkg.in/testfixtures.v2"
	"gochapter/config"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	database, err := sql.Open("mysql", config.Parameters.GetDsn())

	if err != nil {
		log.Fatal(err)
	}

	testfixtures.SkipDatabaseNameCheck(true)

	fixtures, err := testfixtures.NewFolder(database, &testfixtures.MySQL{}, "fixtures")

	if err != nil {
		log.Fatal(err)
	}

	err = fixtures.Load();

	if (err != nil) {
		log.Fatal(err)
	}
}