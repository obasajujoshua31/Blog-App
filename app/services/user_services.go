package services

import (
	"database/sql"
)

type User struct {
	ID         int
	Name       string
	Age        string
	Profession string
	Friendly   bool
	Blogs 		[]Blog
}

const (
	selectByName  = "SELECT * FROM users JOIN blog ON blog.author_id = users.id WHERE users.name=$1;"
	selectByID    = "SELECT * FROM users JOIN blog ON blog.author_id = users.id WHERE users.id=$1;"
	createUserSQL = "INSERT INTO users(name, age, profession, friendly) VALUES ($1, $2, $3, $4) RETURNING *;"
	selectAllUsers = "SELECT * FROM users JOIN blog ON blog.author_id = users.id"
)

func (d *DB) GetUserByName(name string) ([]User, error) {

	rows, err := queryDB(d, name, selectByName)

	if err != nil {
		return nil, err
	}

	var user User

	var blog Blog

	var users []User

	for rows.Next() {
		err = decodeUser(rows, &user, &blog)

		if err != nil {
			return nil, err
		}

		users = append(users, constructUser(user, blog))
	}

	return users, nil
}

func (d *DB) GetAllUsers() ([]User, error) {

	rows, err := queryDB(d, nil, selectAllUsers)

	if err != nil {
		return nil, err
	}

	var user User

	var users []User

	var blog Blog

	for rows.Next() {
		err = decodeUser(rows, &user, &blog)

		if err != nil {
			return nil, err
		}

		users = append(users, constructUser(user, blog))
	}

	return users, nil
}

func (d *DB) GetUserByID(id int) (User, error) {

	rows, err := queryDB(d, id, selectByID)

	if err != nil {
		return User{}, err
	}

	var user User

	var blog Blog

	for rows.Next() {
		err = decodeUser(rows, &user, &blog)

		if err != nil {
			return User{}, err
		}
	}

	retUser := constructUser(user, blog)


	return retUser, nil
}

func (d *DB) CreateUser(user map[string]interface{}) (User, error) {

	stmt, err := d.Prepare(createUserSQL)

	if err != nil {
		return User{}, err
	}

	res, err := stmt.Query(user["name"], user["age"], user["profession"], user["friendly"])

	if err != nil {
		return User{}, err
	}

	var resUser User

	for res.Next() {
		err = res.Scan(&resUser.ID, &resUser.Name, &resUser.Age, &resUser.Profession, &resUser.Friendly)
		if err != nil {
			return User{}, nil
		}
	}

	return resUser, nil
}


func decodeUser(rows *sql.Rows, user *User, blog *Blog) error {
	err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Profession,
		&user.Friendly, &blog.ID, &blog.Title,
		&blog.NumberOfComments, &blog.Content, &blog.AuthorID)
	if err != nil {
		return err
	}

	return nil
}

func constructUser (user User, blog Blog) User {
	var Blogs []Blog

	return User{
		ID:         user.ID,
		Name:       user.Name,
		Age:        user.Age,
		Profession: user.Profession,
		Friendly:   user.Friendly,
		Blogs:      append(Blogs, blog),
	}
}

func queryDB(d *DB, param interface{}, query string) (*sql.Rows, error) {
	stmt, err := d.Prepare(query)

	if err != nil {
		return nil, err
	}

	var rows *sql.Rows

	if param == nil {
		rows, err = stmt.Query()

	} else {
		rows, err = stmt.Query(param)
	}

	if err != nil {
		return nil, err
	}

	return rows, nil
}
