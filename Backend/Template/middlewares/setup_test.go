package middlewares


import (
	"github.com/udaysonu/ober/config"	
	"net/http"
	"testing"
	"os"
)

func Testmain(m *testing.M){
	Repo=NewMiddleware(&config.AppConfig{})
	NewRepo(Repo)
	os.Exit(m.Run())
}


type myHandler struct{

}

func (m *myHandler)ServeHTTP(w http.ResponseWriter,r *http.Request){

}