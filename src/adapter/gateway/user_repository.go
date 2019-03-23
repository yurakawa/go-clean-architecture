package gateway

import (
	"cleanarchitecture/domain"
	"github.com/jinzhu/gorm"
)

type (
	UserRepository struct {
		Conn *gorm.DB
	}

	User struct {
		gorm.Model
		Name  string `gorm:"size:20;not null"`
		Email string `gorm:"size:100;not null"`
		Age   int    `gorm:"type:smallint"`
	}
)

func (r *UserRepository) Store(u domain.User) (id int, err error) {
	user := &User{
		Name:  u.Name,
		Email: u.Email,
		Age: u.Age,
	}

	if err = r.Conn.Create(user).Error; err != nil {
		return
	}

	return int(user.ID), nil
}

func (r *UserRepository) FindByName(name string) (d []domain.User, err error) {
	users := []User{}
	if err = r.Conn.Where("name = ?", name).Find(&users).Error; err != nil {
		return
	}

	n := len(users)
	d = make([]domain.User, n)
	for i := 0; i < n; i++ {
		d[i].ID = int(users[i].ID)
		d[i].Name = users[i].Name
		d[i].Email = users[i].Email
		d[i].Age= users[i].Age
	}
	return
}

func (r *UserRepository) FindAll() (d []domain.User, err error) {
	users := []User{}
	if err = r.Conn.Find(&users).Error; err != nil {
		return
	}

	n := len(users)
	d = make([]domain.User, n)
	for i := 0; i < n; i++ {
		d[i].ID = int(users[i].ID)
		d[i].Name = users[i].Name
		d[i].Email = users[i].Email
		d[i].Age= users[i].Age
	}
	return
}