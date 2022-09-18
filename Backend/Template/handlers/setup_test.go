package handlers

import(
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/udaysonu/ober/config"
	"github.com/udaysonu/ober/helpers"
	mdlr "github.com/udaysonu/ober/middlewares"
	"net/http"
	"time"
)

var session *scs.SessionManager
var appConfig config.AppConfig

func goRoutes()*chi.Mux{
	appConfig.DatabaseName="sql"

	session=scs.New()
	session.Lifetime=24*time.Hour
	session.Cookie.Persist=true
	session.Cookie.SameSite=http.SameSiteLaxMode
	session.Cookie.Secure=false
	appConfig.Session=session

	mux:=chi.NewRouter()
	mux.Use(middleware.Logger)
	mdlrRepo:=mdlr.NewMiddleware(&appConfig)
	mdlr.NewRepo(mdlrRepo)
	mux.Use(mdlrRepo.NoSurf)
	mux.Use(mdlrRepo.CheckMiddleware)
	mux.Use(mdlrRepo.LoadAndSave)

	repo:=NewTestHandler(&appConfig)
	helpers.NewHelpers(&appConfig)

	NewRepo(repo)
	
	mux.Get("/",Repo.Health)
	mux.Get("/register",Repo.AddDriver)
	mux.Get("/drivers",Repo.GetDrivers)
	return mux
}