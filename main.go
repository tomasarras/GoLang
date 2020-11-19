package main

import (
	"api/agencyService"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := StartConn()
	defer db.Close()

	createSchemaIfNotExists(db)
	r := gin.Default()
	r.POST("/agency", postAgencyHandler)
	r.Run()
}

func createSchemaIfNotExists(db *sql.DB) error {
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
	);`

	// execute a query on the server
	_, err := db.Exec(schemaAgency)

	if err != nil {
		return err
	}

	_, err = db.Exec(schemaFlight)
	return err
}

func StartConn() *sql.DB {
	db, err := sql.Open("mysql", "root:password@/flights")
	if err != nil {
		panic(err.Error())
	}

	return db
}

func postAgencyHandler(c *gin.Context) {
	var reqBody agencyService.Agency
	err := c.ShouldBindJSON(&reqBody)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "invalid request body",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"agency": "agencias",
	})
}
