package services

import (
	"time"
)

type User struct {
	ID         int  `gorm:"PRIMARY_KEY; UNIQUE; AUTO_INCREMENT; NOT NULL;" json:"id"`
	Name       string `gorm:"type:varchar(100); NOT NULL;" json:"name"`
	Age        int
	Profession string `gorm:"type:varchar(100); NOT NULL;" json:"profession"`
	Friendly   bool
	Blogs 		[]Blog  `gorm:"foreignkey:AUTHORID;" json:"blogs"`
	CreatedAt time.Time `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp" json:"updated_at"`
}


func (d *DB) GetUserByName(name string) (users []User, err error) {

	err = d.Debug().Preload("Blogs").Where("name = ?", name).Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (d *DB) GetAllUsers() (users []User, err error) {

	err = d.Debug().Preload("Blogs").Find(&users).Error


	if err != nil {
		return nil, err
	}

	return users, nil

}

func (d *DB) GetUserByID(id int) (user User, err error) {

	err = d.Debug().Preload("Blogs").Where("id = ?", id).First(&user).Error


	if err != nil {
		return User{}, err
	}

	return user, nil

}

func (d *DB) CreateUser(args map[string]interface{}) (user User, err error) {

	 user.Name = args["name"].(string)
	 user.Age = args["age"].(int)
	 user.Profession = args["profession"].(string)
	 user.Friendly = args["friendly"].(bool)

	err = d.Debug().Model(&User{}).Create(&user).Error

	if err != nil {
		return User{}, err
	}

	return user, nil
}

