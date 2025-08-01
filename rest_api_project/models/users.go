package models

import (
	"com.shreyash/rest-api/db"
	"com.shreyash/rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := `INSERT INTO users(email,password) VALUES(?,?)`
	stmt,err := db.DB.Prepare(query)
	if err != nil{
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil{
		return err
	}

	res,err := stmt.Exec(u.Email,hashedPassword)
	if err != nil{
		return err
	}

	userId,err := res.LastInsertId()
	u.ID = userId

	return err
}
