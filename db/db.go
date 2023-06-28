package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
)

type DB struct {
	// sql.DB is a pool of 0 or more database conn
	// when we call sql.Open('driver', 'SDN') we will got a *sql.DB
	DB *sql.DB
}

var dbConn = &DB{DB: nil}

// max number of open connections in the db pool
const maxOpenDbConn = 10

// the max number of connections that can remain open but unused in the conn pool
// if a conn is an idle and this limit is reached, this conn will be closed automatically
const maxIdleDbConn = 5

// the maximum time that a specific conn exists (used or unused) and after this time this conn
// will be destroyed even if its used
const maxDbLifeTime = 5 * time.Minute

// function to test the connection before applying the connection
func testDB(db *sql.DB) error {
	// ping the database instance first and check if there is a response or an error returned
	err := db.Ping()
	if err != nil {
		log.Printf("cannot ping the postgres instance \n ERROR ➜ %v", err)
		return err
	}

	// if no error, we recieved the pong
	log.Println("we ping, the postgres instance pong successfully :D")
	return nil // return no errors
}

// function to perform the actuall connection
func ConnectToPostgresInstance(DSN string) (*DB, error) {
	// open a pool of connections
	pool_of_conn, err := sql.Open("pgx", DSN)
	if err != nil {
		log.Printf("cannot open a connection \n ERROR ➜ %v", err)
		return nil, err
	}
	// set pool conn attributes
	pool_of_conn.SetMaxOpenConns(maxOpenDbConn)
	pool_of_conn.SetMaxIdleConns(maxIdleDbConn)
	pool_of_conn.SetConnMaxLifetime(maxDbLifeTime)

	// ping to the connection
	ping_conn_err := testDB(pool_of_conn)
	if ping_conn_err != nil {
		return nil, ping_conn_err
	}
	// set the pool of our returned connection struct instance
	(*dbConn).DB = pool_of_conn
	// return the response
	return dbConn, nil
}
