package services

import (
	"errors"
	"log"
	"sushkof/internal/database"
	"sushkof/internal/models"

	"golang.org/x/crypto/bcrypt"
)

var (
	userRepo database.UserRepo
	taskRepo database.TaskRepo
	not_found = errors.New("invalid email or password")
)

var CookiesMap = make(map[string]*models.User)

func InitRepositories(user_repo database.UserRepo, task_repo database.TaskRepo) {
	userRepo = user_repo
	taskRepo = task_repo
}

func RegisterUser(user models.User) error {
	return CreateUser(user)
}

func Authentication(email, password string) (*models.User, error) {
	user, err := userRepo.FindActiveUserByEmail(email)
	if err != nil {
    log.Println("err in findActiveUser: ", err)
		return nil, not_found
	}
  // пока не подойдет, тк незашифрованные храню
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.Println("Неправильный пароль")
		return nil, not_found
	}
	return &user, nil
}
