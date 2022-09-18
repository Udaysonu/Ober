package dbrepo

import ("github.com/udaysonu/ober/models"
	"time"
	"fmt"
	"log"
	"database/sql"
	_ "github.com/jackc/pgx/v4/stdlib"

)

func (m *PostgresRepo) InsertDriver(driver models.Driver){
	query := `insert into drivers (driver_id,first_name,email,age,phone_number,password,created_at,updated_at) values ($1,$2,$3,$4,$5,$6,$7,$8)`
	_, err := m.SQL.Exec(query, driver.DriverId,driver.FirstName, driver.Email,driver.Age,driver.PhoneNumber,driver.Password, time.Now(), time.Now())
	if err!=nil{
		fmt.Println(err)
	}
}

func (m *PostgresRepo) AllUsers(){
	getAllRows(m.SQL)
}

func getAllRows(conn *sql.DB) error {
	rows, err := conn.Query("select driver_id, first_name, last_name, email, password, age, phone_number from drivers")
	if err != nil {
		log.Println(err)
		return err
	}
	defer rows.Close()
	var firstName, lastName,email,password string
	var id,age,phoneNumber int

	for rows.Next() {
		err := rows.Scan(&id, &firstName, &lastName,&email,&password,&age,&phoneNumber)
		if err != nil {
			log.Println(err)
			return err
		}
		fmt.Println("Record is", id, firstName, lastName)
	}
	if err = rows.Err(); err != nil {
		log.Fatal("Error scanning rows", err)
	}
	fmt.Println("----------------------------")
	return nil
}
