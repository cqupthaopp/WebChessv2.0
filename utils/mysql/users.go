package mysqlutil

/**
 * @Author: Hao_pp
 * @Data: 2022年7月17日-16点03分
 * @Desc: 无无无
 */

import (
	"errors"
	"fmt"
)

func FindUser(name string) error {

	db := Connect(UsersTable)

	defer db.Close()

	var user User

	if _ = db.Where("username = ?", name).Find(&user); user.Username == "" {
		return errors.New("No user called " + name)
	}

	return nil

}

func GetUser(name string) User {

	db := Connect(UsersTable)

	defer db.Close()

	var user User

	db.Where("username = ?", name).Find(&user)

	return user

}

func AddUser(user User) error {

	db := Connect(UsersTable)

	defer db.Close()

	err := db.Create(user)

	fmt.Println(err)

	return nil

}
