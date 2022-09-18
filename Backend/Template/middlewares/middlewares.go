package middlewares

//github.com/udaysonu/ober

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
	"github.com/udaysonu/ober/config"
)

var Repo *MiddlewareRepo

type MiddlewareRepo struct{
	Appconfig *config.AppConfig
}
func NewMiddleware(app *config.AppConfig)(*MiddlewareRepo){
	return &MiddlewareRepo{Appconfig: app}
}
func NewRepo(r *MiddlewareRepo){
	Repo=r
}
func (m *MiddlewareRepo)CheckMiddleware(next http.Handler)http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
		fmt.Println("Middleware tested succesfully....")
		next.ServeHTTP(w,r)
	})
}

func (m *MiddlewareRepo)NoSurf(next http.Handler)http.Handler{
	csrfHandler:=nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:"/",
		Secure:false,
		SameSite:http.SameSiteLaxMode,
	})
	return csrfHandler
}

func (m *MiddlewareRepo)LoadAndSave(next http.Handler)http.Handler{
	return Repo.Appconfig.Session.LoadAndSave(next)
}

// How to use sessions
// 	m.AppConfig.Session.Put(r.Context(),"check-session","working")
// sessionInfo:=m.AppConfig.Session.Get(r.Context(),"check-session")
