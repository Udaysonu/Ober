package dbrepo

import (
	"database/sql"
	"github.com/udaysonu/ober/config"
	"github.com/udaysonu/ober/repository"
)

type PostgresRepo struct {
	SQL *sql.DB
	App *config.AppConfig
}

type TestDBRepo struct{
	DB *sql.DB
	App *config.AppConfig
}


func NewPostgresRepo(sql *sql.DB,app *config.AppConfig)repository.DbRepository{
	return &PostgresRepo{
		SQL:sql,
		App:app,
	}
}

func NewTestDBRepo(app *config.AppConfig)repository.DbRepository{
	return &TestDBRepo{
 		App:app,
	}
}