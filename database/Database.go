package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // don't forget to add it. It doesn't be added automatically
	"log"
	"os"
	"strconv"
	"time"
)

var Db *sql.DB

var (
	ErrUniqueConstraint = errors.New("unique constraint violation")
)

func ConnectDatabase() {

	err := godotenv.Load() //by default, it is .env so we don't have to write
	if err != nil {
		fmt.Println("Error is occurred  on .env file please check")
	}
	//we read our .env file
	host := os.Getenv("HOST")
	port, _ := strconv.Atoi(os.Getenv("PORT")) // don't forget to convert int since port is int type.
	user := os.Getenv("DBUSER")
	dbname := os.Getenv("DB_NAME")
	pass := os.Getenv("PASSWORD")

	// set up postgres sql to open it.
	psqlSetup := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, pass)
	db, errSql := sql.Open("postgres", psqlSetup)
	if errSql != nil {
		log.Println("There is an error while connecting to the database ", err)
		panic(err)
	} else {
		connectionLifetime, _ := strconv.Atoi(os.Getenv("DB_CONNECTION_LIFETIME_SECS"))
		maxConnections, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
		maxIdleConnections, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))

		db.SetConnMaxLifetime(time.Second * time.Duration(connectionLifetime))
		db.SetMaxOpenConns(maxConnections)
		db.SetMaxIdleConns(maxIdleConnections)
		Db = db
		log.Println("Successfully connected to database!")
	}
}
