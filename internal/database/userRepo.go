package database

import (
	"database/sql"
	"log"
	"sushkof/internal/models"
)

type UserRepo interface {
	FindActiveUserByEmail(email string) (models.User, error)
	FindAllActiveUsers() []models.User
	CreateUser(user models.User) error
	DeleteUser(email string) error
	UpdateUserByAdmin(user models.User) error
	UpdateUserByUser(user models.User) error
}

type UserRepoDB struct {
	*sql.DB
}

func (udb UserRepoDB) FindActiveUserByEmail(email string) (models.User, error) {
	user := models.User{}
	row := udb.DB.QueryRow("select name, age, email, password, admin from users where email = ? and active = true", email)
	err := row.Scan(&user.Name, &user.Age, &user.Email, &user.Password, &user.Admin)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (udb UserRepoDB) FindAllActiveUsers() []models.User {
	var users []models.User
	rows, err := udb.DB.Query("select name, age, email, password, admin from users where active = true order by name")
	if err != nil {
		log.Println("err in query: ", err)
		return users
	}

	for rows.Next() {
		user := models.User{}
		err := rows.Scan(&user.Name, &user.Age, &user.Email, &user.Password, &user.Admin)
		if err != nil {
			log.Println("err in scan: ", err)
			return users
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		log.Println("err in Err: ", err)
		return users
	}
	return users
}

func (udb UserRepoDB) CreateUser(user models.User) error {
	_, err := udb.DB.Exec("insert into users(name, age, email, password, admin) values(?,?,?,?,?)",
		user.Name, user.Age, user.Email, user.Password, user.Admin)

	return err
}

func (udb UserRepoDB) DeleteUser(email string) error {
	_, err := udb.Exec("update users set active = false where email = ?", email)
	return err
}

func (udb UserRepoDB) UpdateUserByAdmin(user models.User) error {
	_, err := udb.Exec("update users set name = ?, password = ?, age = ?, admin = ? where email = ?",
		user.Name, user.Password, user.Age, user.Admin, user.Email)
	return err
}

func (udb UserRepoDB) UpdateUserByUser(user models.User) error {
	_, err := udb.Exec("update users set name = ?, password = ?, age = ? where email = ?",
		user.Name, user.Password, user.Age, user.Email)
	return err
}
