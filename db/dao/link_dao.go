package dao

import (
    "database/sql"
    _ "github.com/lib/pq"
    "sync"
)

type Link struct {
    Token       string `db:"token"`
    CreatedById int    `db:"created_by_id"`
    ExpiresAt   *int   `db:"expires_at"`
    Link        string `db:"link"`
}

//postgresql://postgres:short@db:5432/short

type LinkDAO struct {
    db *sql.DB
}

func (l *LinkDAO) getDB() *sql.DB {
    once := sync.Once{}

    once.Do(func() {
        var err error
        l.db, err = sql.Open("postgres", "postgres://postgres:short@localhost:5432/short?sslmode=disable")
        if err != nil {
            panic(err)
        }
    })

    //error ping
    if l.db.Ping() != nil {

    }

    return l.db
}

func (l *LinkDAO) Create(link Link) error {
    _, err := l.getDB().Exec(
        "INSERT INTO links (token, created_by_id, link) VALUES ($1, $2, $3)",
        link.Token,
        link.CreatedById,
        link.Link,
    )

    if err != nil {
        return err
    }

    return nil
}

func (l *LinkDAO) GetByShort(short string) *Link {
    link := Link{}

    l.getDB().QueryRow(
        "SELECT token, created_by_id, link FROM links WHERE token = $1 LIMIT 1",
        short,
    ).Scan(
        &link.Token,
        &link.CreatedById,
        &link.Link,
    )

    return &link
}

func (l *LinkDAO) GetByLink(link string) *Link {
    linkEntity := Link{}

    l.getDB().QueryRow(
        "SELECT token, created_by_id, link FROM links WHERE link = $1 LIMIT 1",
        link,
    ).Scan(
        &linkEntity.Token,
        &linkEntity.CreatedById,
        &linkEntity.Link,
    )

    if linkEntity.Token == "" {
        return nil
    }

    return &linkEntity
}
