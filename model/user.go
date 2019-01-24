package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
}

func CheckUser(username string, password string) (string, bool) {
	var user User

	err := db.Where(User{Username: username}).First(&user).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return "用户名不存在", false
	} else if err != nil {
		return err.Error(), false
	}

	if user.Password != password {
		return "密码错误", false
	} else {
		return "校验成功", true
	}
}

func FindUserByUsername(username string) User {
	var user User
	db.Where(User{Username: username}).First(&user)
	return user
}

func CreateNewUser(username, password string) string {
	user := User{Username: username, Password: password}
	err := db.Create(&user).Error
	if err == nil {
		return "创建完成"
	} else {
		return "创建失败"
	}		
}

func ChangeUserPassword(username, password string) string {
	var user User
	err := db.Model(&user).Where("username = ?", username).Update("password", password).Error
	if err != nil {
		return err.Error()
	} else {
		return "OK"
	}
} 
 
func DeleteUserByUsername(username string) string {
	var user User
	err := db.Where("username = ?", username).First(&user).Delete(&user).Error
	if err != nil {
		return err.Error()
	} else {
		return "OK"
	}
}