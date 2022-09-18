package repository

import ("github.com/udaysonu/ober/models")


type DbRepository interface{
	AllUsers()
	InsertDriver(driver models.Driver)
}