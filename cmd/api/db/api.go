package db

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


type User struct {
	Uid      string 
	Name     string
	Password string
}


func GetUserByUsername(name string) (*User, error) {
	user := User{}
	if err := DB.Where("name = ?", name).First(&user).Error; err != nil {
		log.Error(err)
		return user, err
    }
	return user, nil
}
