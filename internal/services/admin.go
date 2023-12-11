package services

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
	"sushkof/internal/models"
)

func CreateUser(user models.User) error {
	if _, err := userRepo.FindActiveUserByEmail(user.Email); err != nil {
		//  надо пароль проверить на sql инъекцию, либо зашифровать
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err // там возвращается ошибка, что пароль длинее 72 символов
			// можно самому проверить и убрать это
		}

		user.Password = string(hashedPassword)

		err = userRepo.CreateUser(user)
		if err != nil {
			// ошибка при занесении записи в бд
			log.Println("error in create: ", err)
			return err
		}
		// если условие не сработало, то зарегались и возвращаем что надо
		return nil
	}
	// пользователь уже существует в бд
	return errors.New("user is already exist")
}

func DeleteUser(email string) error {
	if _, err := userRepo.FindActiveUserByEmail(email); err == nil {
		err := userRepo.DeleteUser(email)
		if err != nil {
			// ошибка при занесении записи в бд
			log.Println("error in delete: ", err)
			return err
		}
		// если условие не сработало, то "удалили"
		return nil
	}
	// такого пользователя нет, поэтому не можем удалить
	return not_found
}

type TaskAndError struct {
	Error string      `json:"error"`
	Task  models.Task `json:"task"`
}

// Функция для созданий заданий любым пользователям
func CreateTaskForMultipleUsers(task models.Task, users []string) []TaskAndError {
	taskAndErrors := make([]TaskAndError, len(users))
  // var isError bool
	for i, email := range users {
    var tae TaskAndError
		task.User_email = email
		task, err := taskRepo.CreateTask(task); 
    if err != nil {
      // isError = true // один бы раз сделать, а не каждую ошибку
		  tae.Error = err.Error()
    }
    tae.Task = task
		taskAndErrors[i] = tae
	}
	return taskAndErrors//, isError
}
