
package services

import (
	"database/sql"
)

type Blog struct {
	ID         int
	Title       string
	NumberOfComments        int
	Content string
	AuthorID   int
	User 	User
}

const (
	selectBlogByID    = "SELECT * FROM blog JOIN users ON users.id = blog.author_id WHERE blog.id=$1;"
	createBlogSQL = "INSERT INTO blog(title, content, author_id) VALUES ($1, $2, $3) RETURNING *;"
	selectAllBlogs = "SELECT * FROM blog JOIN users ON users.id = blog.author_id;"
)


func (d *DB) GetAllBlogs() ([]Blog, error) {

	rows, err := queryDB(d, nil, selectAllBlogs)

	if err != nil {
		return nil, err
	}

	var blog Blog

	var blogs []Blog

	for rows.Next() {
		err = decodeBlog(rows, &blog)

		if err != nil {
			return nil, err
		}

		blogs = append(blogs, blog)
	}

	return blogs, nil
}

func (d *DB) GetBlogByID(id int) (Blog, error) {

	rows, err := queryDB(d, id, selectBlogByID)

	if err != nil {
		return Blog{}, err
	}

	var blog Blog

	for rows.Next() {
		err = decodeBlog(rows, &blog)

		if err != nil {
			return Blog{}, err
		}
	}

	return blog, nil
}

func (d *DB) CreateBlog(blog map[string]interface{}) (Blog, error) {

	stmt, err := d.Prepare(createBlogSQL)

	if err != nil {
		return Blog{}, err
	}

	res, err := stmt.Query(blog["title"], blog["content"], blog["author_id"])

	if err != nil {
		return Blog{}, err
	}

	var resBlog Blog

	for res.Next() {
		err = res.Scan(&resBlog.ID, &resBlog.Title, &resBlog.NumberOfComments, &resBlog.Title, &resBlog.AuthorID)
		if err != nil {
			return Blog{}, err
		}
	}

	stmt, err = d.Prepare(selectBlogByID)

	if err != nil {
		return Blog{}, nil
	}

	results, err := stmt.Query(resBlog.ID)

	var newBlog Blog

	if err != nil {
		return Blog{}, nil
	}

	for results.Next() {
		err = decodeBlog(results, &newBlog)
		if err != nil {
			return Blog{}, nil
		}
	}

    return  newBlog, nil
}

func decodeBlog(rows *sql.Rows, blog *Blog) error {
	err := rows.Scan(&blog.ID, &blog.Title, &blog.NumberOfComments, &blog.Content, &blog.AuthorID, &blog.User.ID,
		&blog.User.Name, &blog.User.Age, &blog.User.Profession, &blog.User.Friendly)

	if err != nil {
		return err
	}

	return nil
}

