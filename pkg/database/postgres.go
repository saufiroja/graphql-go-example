package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/saufiroja/graphql-go-example/config"
	"github.com/sirupsen/logrus"
)

type IPostgres interface {
	Db() *sql.DB
	StartTransaction() (*sql.Tx, error)
	CommitTransaction(tx *sql.Tx)
	RollbackTransaction(tx *sql.Tx)
}

type Postgres struct {
	*sql.DB
}

func NewPostgres(conf *config.AppConfig) IPostgres {
	host := conf.Postgres.Host
	port := conf.Postgres.Port
	user := conf.Postgres.User
	pass := conf.Postgres.Pass
	dbname := conf.Postgres.Name
	ssl := conf.Postgres.SSL

	str := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, pass, dbname, ssl)

	db, err := sql.Open("postgres", str)
	if err != nil {
		panic(err)
	}

	// check connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// set max connection
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5)
	db.SetConnMaxIdleTime(5)

	logrus.Info("Connected to database")

	return &Postgres{db}
}

func (p *Postgres) Db() *sql.DB {
	return p.DB
}

func (p *Postgres) StartTransaction() (*sql.Tx, error) {
	return p.Begin()
}

func (p *Postgres) CommitTransaction(tx *sql.Tx) {
	err := tx.Commit()
	if err != nil {
		panic(err)
	}
}

func (p *Postgres) RollbackTransaction(tx *sql.Tx) {
	err := tx.Rollback()
	if err != nil {
		panic(err)
	}
}

func (p *Postgres) Close(ctx context.Context) {
	err := p.DB.Close()
	if err != nil {
		panic(err)
	}
}
