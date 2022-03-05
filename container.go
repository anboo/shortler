package main

import (
    "database/sql"
    "fmt"
    "os"
    "sync"
)

type Container struct {
    db *sql.DB
}

func (l *Container) getDB() *sql.DB {
    once := sync.Once{}

    connect := func() {
        var err error
        l.db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
        if err != nil {
            panic(err)
        }
    }

    once.Do(func() {
        connect()
    })

    pingErr := l.db.Ping()
    if pingErr != nil {
        fmt.Println("error db ping: " + pingErr.Error())
        connect()
    }

    return l.db
}
