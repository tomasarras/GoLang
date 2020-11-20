package database

import (
	"database/sql"
)

func StartConn() *sql.DB {
	db, err := sql.Open("mysql", "root:password@/flights")
	if err != nil {
		panic(err.Error())
	}

	//createDatabaseIfNotExist(db)
	createSchemaIfNotExists(db)
	return db
}

func createDatabaseIfNotExist(db *sql.DB) {
	s := "CREATE DATABASE IF NOT EXISTS flights;"
	_, err := db.Exec(s)
	if err != nil {
		panic(err.Error())
	}

	s = "USE flights;"
	_, err = db.Exec(s)
	if err != nil {
		panic(err.Error())
	}

	createSchemaIfNotExists(db)
}

func createSchemaIfNotExists(db *sql.DB) {
	schemaAgency := `CREATE TABLE IF NOT EXISTS agency (
		id_agency int NOT NULL AUTO_INCREMENT,
		name varchar(50) NOT NULL,
		CONSTRAINT PK_AGENCY PRIMARY KEY (id_agency)
	);`

	schemaFlight := `CREATE TABLE IF NOT EXISTS flight (
		id_flight int NOT NULL AUTO_INCREMENT,
		name varchar(50) NOT NULL,
		start timestamp NOT NULL,
		end timestamp NOT NULL,
		aircraft varchar(50) NOT NULL,
		id_agency int NOT NULL,
		CONSTRAINT PK_FLIGHT PRIMARY KEY (id_flight),
		FOREIGN KEY FK_FLIGHT_AGENCY (id_agency)
		REFERENCES agency (id_agency)
		ON DELETE CASCADE
	);`

	// execute a query on the server
	_, err := db.Exec(schemaAgency)

	if err != nil {
		panic(err.Error())
	}

	_, err = db.Exec(schemaFlight)
	if err != nil {
		panic(err.Error())
	}
}
