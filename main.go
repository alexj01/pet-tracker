package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func main() {

	var (
		DBNAME       = os.Getenv("dbName")
		DBHOST       = os.Getenv("dbHost")
		DBPORT       = 5432
		DBUSER       = os.Getenv("dbUser")
		DBPASS       = os.Getenv("dbPassword")
		DBCERTPATH   = os.Getenv("dbCerts")
		DBCLIENTCERT = "/client-cert.pem"
		DBROOTCERT   = "/server-ca.pem"
		DBSSLKEY     = "/client-key.pem"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=verify-ca sslcert=%s sslkey=%s sslrootcert=%s",
		DBHOST, DBPORT, DBUSER, DBPASS, DBNAME, DBCERTPATH+DBCLIENTCERT, DBCERTPATH+DBSSLKEY,
		DBCERTPATH+DBROOTCERT)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}
