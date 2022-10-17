package handlers

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/albrow/forms"
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

func (m *Repository) PostTest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
    var t models.Driver
    err := decoder.Decode(&t)
    if err != nil {
        panic(err)
    }
    fmt.Println(t)


	// sending response in json format
	out,_:=json.MarshalIndent(t,"","  ")
	w.Header().Set("Content-Type","application/json")

	w.Write(out)
	return
}
func (m *Repository) CheckSession(w http.ResponseWriter, r *http.Request){
	res,ok:=m.AppConfig.Session.Get(r.Context(),"session-data").(models.Driver)
	fmt.Println(res,ok)
	fmt.Println(m.AppConfig.Session.Get(r.Context(),"session-data"))
}

func (m *Repository) CheckPostRequest(w http.ResponseWriter, r *http.Request){
	 	// Parse request data.

		fmt.Println("&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&",json.NewDecoder(r.Body))
	userData, err := forms.Parse(r)

	if err != nil {
		fmt.Println(err)
		// Handle err
		// ...
	}
	fmt.Println(userData.Values,err)
 	// Validate
	val := userData.Validator()
	fmt.Println(userData.Get("first_name"))
	val.Require("first_name")
	val.LengthRange("first_name", 4, 16)
	val.Require("email")
	val.MatchEmail("email")
	val.Require("password")
	val.MinLength("password", 8)
	// val.Require("confirmPassword")
	// val.Equal("password", "confirmPassword")
 	if val.HasErrors() {
		 fmt.Println(val.ErrorMap())
 	}
 
 
}