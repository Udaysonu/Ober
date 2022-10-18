package main
import (
	"github.com/joho/godotenv"
	"github.com/go-chi/chi/v5"
	"github.com/udaysonu/ober/config"
	"github.com/udaysonu/ober/driver"
	"testing"
	"fmt"
	"os"
)

var DB *driver.DB

// func TestRun(t *testing.T){
// 	godotenv.Load("../../.env")
//  	dbHostname:=os.Getenv("DB_HOSTNAME")
// 	dbPort:=os.Getenv("DB_PORT")
// 	dbName:=os.Getenv("DB_NAME")
// 	dbUser:=os.Getenv("DB_USER")
// 	dbPassword:=os.Getenv("DB_PASSWORD")
// 	dsn:=fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s",dbHostname,dbPort,dbName,dbUser,dbPassword)

// 	db,err:=Run(dsn)
// 	DB=db
// 	if err!=nil{
// 		t.Error("Error while running Run")
// 	}
// }

// func TestRoutes(t *testing.T){
// 	var app config.AppConfig;
//  	mux:=routes(&app,DB)
// 	switch val:=mux.(type){
// 	case *chi.Mux:
// 	default: 
// 		t.Errorf("Routes Expected: *chi.Mux, Got:%T",val)
// 	}
// }