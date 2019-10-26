package services

import (
	"database/sql"
)

type User struct {
	ID int
	Name string
	Age string
	Profession string
	Friendly bool
}

const (
	selectByName = "SELECT * FROM users WHERE name=$1"
	selectByID = "SELECT * FROM users WHERE id=$1"
	)

func (d *DB) GetUserByName(name string) ([]User, error) {

	rows, err := queryDB(d, name, selectByName)

	if err != nil {
		return nil, err
	}

	var user User

	var users []User

	for rows.Next() {
		err = decodeUser(rows, &user)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (d *DB) GetUserByID(id int) (User, error) {

	rows, err := queryDB(d, id, selectByID)

	if err != nil {
		return User{}, err
	}

	var user User

	err = decodeUser(rows, &user)

	if err != nil {
		return User{}, err
	}
	return user, nil
}

func decodeUser(rows *sql.Rows, user *User) error{
	err:= rows.Scan(&user.ID, &user.Name, &user.Age, &user.Profession, &user.Friendly)

	if err != nil {
		return err
	}

	return nil
}

func queryDB(d *DB, param interface{}, query string)  (*sql.Rows, error,) {
	stmt, err := d.Prepare(query)

	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(param)

	if err != nil {
		return nil, err
	}

	return rows, nil
}