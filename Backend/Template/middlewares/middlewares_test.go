package middlewares

import (
	"github.com/udaysonu/ober/config"
	"net/http"
	"testing"
)

func TestCheckMiddleware(t *testing.T){
	myH:=myHandler{}
	response:=Repo.CheckMiddleware(&myH)
	switch val:=response.(type){
	case http.Handler:
	default:
		t.Errorf("CheckMiddleware Error: Expected: http.Handler, Got: %T",val)
	}
}

func TestNoSurf(t *testing.T){
	myH:=myHandler{}
	response:=Repo.NoSurf(&myH)
	switch val:=response.(type){
	case http.Handler:
	default :
		t.Errorf("NoSurf Error: Expected: http.Handler, Got: %T",val)
	}
}

func TestLoadAndSave(t *testing.T){
	NewRepo(NewMiddleware(&config.AppConfig{}))
	myH:=myHandler{}
	response:=Repo.LoadAndSave(&myH)
	switch val:=response.(type){
	case http.Handler:
	default :
		t.Errorf("LoadAndSave Error: Expected: http.Handler, Got: %T",val)
	}

}