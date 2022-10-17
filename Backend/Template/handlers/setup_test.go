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
	"testing"
	"os"
)

var session *scs.SessionManager
var appConfig config.AppConfig

func TestMain(m *testing.M){
	appConfig.DatabaseName="sql"

	session=scs.New()
	session.Lifetime=24*time.Hour
	session.Cookie.Persist=true
	session.Cookie.SameSite=http.SameSiteLaxMode
	session.Cookie.Secure=false
	appConfig.Session=session

	repo:=NewTestHandler(&appConfig)
	helpers.NewHelpers(&appConfig)

	NewRepo(repo)
	
	os.Exit(m.Run())
}

func goRoutes()*chi.Mux{

	mux:=chi.NewRouter()
	mux.Use(middleware.Logger)
	mdlrRepo:=mdlr.NewMiddleware(&appConfig)
	mdlr.NewRepo(mdlrRepo)
	mux.Use(mdlrRepo.NoSurf)
	mux.Use(mdlrRepo.CheckMiddleware)
	mux.Use(mdlrRepo.LoadAndSave)

	mux.Get("/",Repo.Health)
	mux.Post("/register",Repo.AddDriver)
	mux.Get("/drivers",Repo.GetDrivers)
	return mux
}