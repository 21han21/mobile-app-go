package database

import (
    "database/sql"
    _ "github.com/lib/pq"
    "log"
    "fmt"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "solo2121"
    dbname   = "informingdb"
)

func InitDB() (*sql.DB, error) {
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)
    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        return nil, err
    }
    if err = db.Ping(); err != nil {
        return nil, err
    }
    log.Println("Successfully connected to the database!")
    return db, nil
}
