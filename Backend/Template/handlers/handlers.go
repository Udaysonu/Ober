package handlers

import (
	"fmt"
	"net/http"

	"github.com/udaysonu/ober/config"
	"github.com/udaysonu/ober/driver"
	"github.com/udaysonu/ober/models"
	"github.com/udaysonu/ober/repository"
	"github.com/udaysonu/ober/repository/dbrepo"
)

var Repo *Repository

type Repository struct {
	AppConfig *config.AppConfig
	SqlDB     repository.DbRepository
}

func NewRepo(r *Repository) {
	Repo = r
}

 

func NewHandler(appconfig *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{AppConfig: appconfig, SqlDB: dbrepo.NewPostgresRepo(db.SQL, appconfig)}
}

func NewTestHandler(appconfig *config.AppConfig) *Repository {
	return &Repository{AppConfig: appconfig, SqlDB: dbrepo.NewTestDBRepo(appconfig)}
}

func (m *Repository) Health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "server_status:up")
}

func (m *Repository) AddDriver(w http.ResponseWriter, r *http.Request) {
	m.SqlDB.InsertDriver(models.Driver{DriverId:"12345",FirstName:"uday",Email:"uday@gmail.com",Age:20,PhoneNumber:9889898,Password:"udayyadusonu@1"})
}


func (m *Repository) GetDrivers(w http.ResponseWriter, r *http.Request) {
	m.SqlDB.AllUsers()
}
