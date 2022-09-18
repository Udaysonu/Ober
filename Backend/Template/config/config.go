package config

import (
	"log"
	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	DatabaseName string
	Session      *scs.SessionManager
	InfoLog		 *log.Logger
	ErrorLog 	 *log.Logger
	InProduction bool
}