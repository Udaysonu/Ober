package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"github.com/alexedwards/scs/v2"
	"github.com/joho/godotenv"
	"github.com/udaysonu/ober/config"
	"github.com/udaysonu/ober/driver"
)


var session *scs.SessionManager
var appConfig config.AppConfig

func main(){
	appConfig.DatabaseName="sql"
	appConfig.InfoLog=log.New(os.Stdout,"INFO\t",log.Ldate|log.Ltime)
	appConfig.ErrorLog=log.New(os.Stdout,"ERROR\t",log.Ldate|log.Ltime|log.Lshortfile)
	appConfig.InProduction=false

	err1:= godotenv.Load("../../.env")
	if err1!=nil{
		log.Fatal("Error loading .env file")
	}

	serverPort:=":"+os.Getenv("SERVER_PORT")
 	dbHostname:=os.Getenv("DB_HOSTNAME")
	dbPort:=os.Getenv("DB_PORT")
	dbName:=os.Getenv("DB_NAME")
	dbUser:=os.Getenv("DB_USER")
	dbPassword:=os.Getenv("DB_PASSWORD")
	dsn:=fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s",dbHostname,dbPort,dbName,dbUser,dbPassword)

	db,err:=Run(dsn)

	if(err!=nil){
		log.Println("Unable to connect to the server")
	}

	defer db.SQL.Close()

	fmt.Println("server starting....")

	err=http.ListenAndServe(serverPort,routes(&appConfig,db))
	if err!=nil{
		fmt.Println("Unable to start server",err)
	}else{
		fmt.Println("server started at port:8080")
	}
}


func Run(msn string)(*driver.DB, error){
	db,err:=driver.ConnectSQL(msn)
	if err!=nil{
		return nil,err
	}
	log.Println("Connected to the Database")
	session=scs.New()
	session.Lifetime=24*time.Hour
	session.Cookie.Persist=true
	session.Cookie.SameSite=http.SameSiteLaxMode
	session.Cookie.Secure=false
	appConfig.Session=session

	return db,nil

}

