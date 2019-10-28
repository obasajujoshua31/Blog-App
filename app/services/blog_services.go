
package services

import "time"

type Blog struct {
	ID        int    `gorm:"primary_key;auto_increment;" json:"id"`
	Title       string `gorm:"type:varchar(100); NOT NULL;" json:"title"`
	NumberOfComments        int  `gorm:"DEFAULT 0;" json:"number_of_comments"`
	Content string  `gorm:"type:varchar(400); NOT NULL;" json:"content"`
	AuthorID   int `json:"author_id"`
	User 	User `gorm:"foreignkey:author_id;" json:"user"`
	CreatedAt time.Time `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp" json:"updated_at"`
}

func (d *DB) GetAllBlogs() (blogs []Blog, err error) {

	err = d.Debug().Preload("User").Find(&blogs).Error

	if err != nil {
		return nil, err
	}

	return blogs, nil
}

func (d *DB) GetBlogByID(id int) (blog Blog, err error) {

	err = d.Debug().Preload("User").Where("id = ?", id).Take(&blog).Error

	if err != nil {
		return Blog{}, err
	}

	return blog, nil
}


func (d *DB) CreateBlog(args map[string]interface{}) (blog Blog, err error) {

	blog.AuthorID = args["author_id"].(int)
	blog.Content = args["content"].(string)
	blog.Title = args["title"].(string)


	err = d.Debug().Model(&Blog{}).Create(&blog).Error

	if err != nil {
		return Blog{}, err
	}

	return blog, nil

}


