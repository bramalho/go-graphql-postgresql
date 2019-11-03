package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Db struct {
	*sql.DB
}

func New(connString string) (*Db, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Db{db}, nil
}

func ConnString(host string, port int, user string, password string, dbName string) string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName,
	)
}

type User struct {
	ID     int
	Name   string
	Age    int
	Status bool
}

func (d *Db) GetUsersByName(name string) []User {
	stmt, err := d.Prepare("SELECT * FROM users WHERE name LIKE $1")
	if err != nil {
		fmt.Println("GetUserByName Preperation Err: ", err)
	}

	rows, err := stmt.Query(name)
	if err != nil {
		fmt.Println("GetUserByName Query Err: ", err)
	}

	var r User
	users := []User{}
	for rows.Next() {
		err = rows.Scan(
			&r.ID,
			&r.Name,
			&r.Age,
			&r.Status,
		)
		if err != nil {
			fmt.Println("Error scanning rows: ", err)
		}
		fmt.Println(r)
		users = append(users, r)
	}

	return users
}
