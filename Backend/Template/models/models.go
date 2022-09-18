package models

type Driver struct{
	DriverId string `json:"driver_id"`
	FirstName string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName string  `json:"last_name"`
	Email string `json:"email"`
	Age uint64 	`json:"age"`
	PhoneNumber uint64 `json:"phone_number"`
	Password string `json:"password"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Bio string  `json:"bio"`
}