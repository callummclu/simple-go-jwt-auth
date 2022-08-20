package models

import (
	"database/sql"
	"errors"
	"go-react-auth/auth"
	"go-react-auth/configs"
)

type User struct {
	ID       int64  `json:"-"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
}

type LogInUser struct {
	User     string `json:"username"`
	Password string `json:"password"`
}

func NewUser() *User {
	return new(User)
}

func NewLogInUser() *LogInUser {
	return new(LogInUser)
}

func (u *User) SaveUser() error {
	db, err := configs.GetDB()

	if err != nil {
		return err
	}

	var (
		username string
	)
	stmt, err := db.Prepare("SELECT username FROM users WHERE username = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(u.Username).Scan(&username)
	if err != nil {
		if err == sql.ErrNoRows {
			hashedPassword, err := auth.HashPassword(u.Password)
			if err != nil {
				return err
			}
			//Add the new user
			insert_stmt, err := db.Prepare("INSERT INTO users (username,password) VALUES ($1,$2)")

			if err != nil {
				return err
			}
			defer insert_stmt.Close()
			_, err = insert_stmt.Exec(u.Username, hashedPassword)

			return err
		} else {
			return err
		}
	} else {
		err = errors.New("User already exists")
		return err
	}
}

func (u *LogInUser) UserLogin() (string, error) {
	db, err := configs.GetDB()
	if err != nil {
		return "", err
	}
	defer db.Close()
	stmt, err := db.Prepare("SELECT username,password FROM users WHERE username = $1")
	if err != nil {
		return "", err
	}
	defer stmt.Close()
	var (
		username string
		password string
	)
	err = stmt.QueryRow(u.User).Scan(&username, &password)

	if err != nil {
		return "", err
	}
	err = auth.CheckPassword(password, u.Password)

	if err != nil {
		return "", err
	}
	return username, err
}
