package main

import (
	"github.com/GuiaBolso/darwin"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
	"gochapter/config"
)

var (
	migrations = []darwin.Migration{
		{
			Version:     1,
			Description: "Creating table cars",
			Script: `CREATE TABLE cars (
				id               INT auto_increment,
				name 		   VARCHAR(255),
				day_max_distance INT,
				type             INT,
				price_per_km     INT,
				rent_price       INT,
				image            TEXT,
				PRIMARY KEY (id)
			) ENGINE=InnoDB CHARACTER SET=utf8;`,
		},
		{
			Version:     2,
			Description: "Update cars, add dates",
			Script: `ALTER TABLE cars
			 	 ADD COLUMN createdAt DATETIME NULL DEFAULT NULL,
			 	 ADD COLUMN updatedAt DATETIME NULL DEFAULT NULL;`,
		},
	}
)

func main() {
	database, err := sql.Open("mysql", config.Parameters.GetDsn())

	if err != nil {
		log.Fatal(err)
	}

	driver := darwin.NewGenericDriver(database, darwin.MySQLDialect{})

	d := darwin.New(driver, migrations, nil)

	err = d.Migrate()

	if err != nil {
		log.Println(err)
	}
}