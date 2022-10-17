package main

import (
	"net/http"
	"github.com/udaysonu/ober/driver"
	"github.com/udaysonu/ober/config"
	mdlr "github.com/udaysonu/ober/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/udaysonu/ober/handlers"
	"github.com/udaysonu/ober/helpers"
)



func routes(app *config.AppConfig, db *driver.DB)http.Handler{
	mux:=chi.NewRouter()
	mux.Use(middleware.Logger)
	mdlrRepo:=mdlr.NewMiddleware(app)
	mdlr.NewRepo(mdlrRepo)
	// mux.Use(mdlrRepo.NoSurf)
	mux.Use(mdlrRepo.CheckMiddleware)
	mux.Use(mdlrRepo.LoadAndSave)

	repo:=handlers.NewHandler(&appConfig,db)
	helpers.NewHelpers(&appConfig)

	handlers.NewRepo(repo)
	mux.Post("/post",handlers.Repo.PostTest)
	mux.Post("/register",handlers.Repo.CheckPostRequest)
	mux.Get("/drivers",handlers.Repo.GetDrivers)
	mux.Get("/",handlers.Repo.Health)

	return mux
}