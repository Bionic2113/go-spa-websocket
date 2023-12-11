package services

import (
	"encoding/json"
	"log"
	"sushkof/internal/models"
)

type Action interface {
	MakeAction(message Message) []byte
}

type ActionFunc func(message Message) []byte

func (a ActionFunc) MakeAction(message Message) []byte {
	return a(message)
}

type Message struct {
	Type string `json:"type"`
	Obj  string `json:"obj"`
}

type Response struct {
  Error bool   `json:"error"`
  Type  string `json:"type"`
  Obj   string `json:"obj,omitempty"`
}

type err_message string

func (e err_message) marshal() []byte{
  bytes, _ := json.Marshal(e)
  return bytes
}

const (
  incorrect_data err_message = "Ошибка: неверные данные"
  email_err err_message = "Ошибка: пользователь с таким email уже существует"
  already_deleted err_message = "Ошибка: уже нет в базе"
  already_exist err_message = "Ошибка: уже существует в базе"
  not_found_err err_message = "Ошибка: не найдено в базе"
  not_found_email err_message = "Ошибка: нет пользователя с таким email"
)

func (r Response) Marshal() []byte{
  bytes, _ := json.Marshal(r)
  return bytes
}

func (r Response) Unmarshal(bytes []byte) (Response, error) {
  err := json.Unmarshal(bytes, &r)
  return r, err
}

var Actions = map[string]Action{
    "current user"                   : actionCurrentUser,
    "create user"                    : ActionCreateUser,
    "delete user"                    : ActionDeleteUser,
    "find user"                      : ActionFindUser,
    "find all users"                 : ActionFindAllUsers,
    "update user by user"            : ActionUpdateUserByUser,
    "update user by admin"           : ActionUpdateUserByAdmin,

    "create task"                    : ActionCreateTask,
    "delete task"                    : ActionDeleteTask,
    "update task"                    : ActionUpdateTask,
    "find all task for user"         : ActionFindAllTaskForUser,
    "create task for multiple users" : ActionCreateTaskForMultipleUsers,
}

/*
    НЕ ВЕЗДЕ НАПИСАЛ МЕТОДЫ В СЕРВИСЕ
    ПОЭТОМУ У ТЕБЯ ЕСТЬ С РЕПОЗИТОРИЕВ
    СКОРЕЕ ВСЕГО ЛУЧШЕ ДОБАВИТЬ В РЕПОЗИТОРИИ МЕТОДЫ
    ТАКЖЕ КУЧА ОДИНАКОВЫХ ПРОВЕРОК
    МОЖЕТ СМОЖЕШЬ КАК ТО ИСПРАВИТЬ

    Нужен рефакторинг!
*/

var actionCurrentUser                ActionFunc = func(message Message) []byte {
  response := Response{Type: message.Type}
  var user_uuid string
  err := json.Unmarshal([]byte(message.Obj), &user_uuid)
  if err != nil { // вообще ненужная тема так то, надо будет убрать после тестов
    log.Println("ошибка в десериализации строки с куками")
    response.Error = true
    response.Obj = string(not_found_err.marshal())
    return response.Marshal()
  }
  user, ok := CookiesMap[user_uuid]
  if !ok {
    log.Println("ошибка в получения юзера в мапе с куками")
    response.Error = true
    response.Obj = string(not_found_err.marshal())
    return response.Marshal()

  }
  response.Obj = string(user.Marshal())
  response.Error = false // need delete
  return response.Marshal()
}

var ActionCreateUser                 ActionFunc = func(message Message) []byte { 
  var user models.User
  response := Response{Type: message.Type}
  user, err := user.Unmarshal([]byte(message.Obj))
  if err != nil{
    log.Println("err in unmarshall user: ", err)
    response.Error = true
    response.Obj = string(incorrect_data.marshal())
    return response.Marshal()
  }
  err = CreateUser(user) // надо чекнуть, что данные не перезатруться при создании с таким же мылом
  if err != nil {
    response.Error = true
    response.Obj = string(email_err.marshal())
    return response.Marshal()
  }
  response.Obj = message.Obj
  // response.Error = false // по идее не надо тк дефолт значение
  return response.Marshal()
}

var ActionDeleteUser                 ActionFunc = func(message Message) []byte { 
  var email string
  response := Response{Type: message.Type}
  err := json.Unmarshal([]byte(message.Obj), &email)
  if err != nil {
    response.Error = true
    response.Obj = string(incorrect_data.marshal())
  }
  err = DeleteUser(email)
  if err != nil {
    response.Error = true
    response.Obj = string(already_deleted .marshal())
  }
  response.Error = false // по идее не надо тк дефолт значение
  response.Obj = message.Obj
  return response.Marshal()
}

var ActionFindUser                   ActionFunc = func(message Message) []byte {
var email string
  response := Response{Type: message.Type}
  err := json.Unmarshal([]byte(message.Obj), &email)
  if err != nil {
    response.Error = true
    response.Obj = string(incorrect_data.marshal())
  }
  user, err := userRepo.FindActiveUserByEmail(email)
  if err != nil {
    response.Error = true
    response.Obj = string(already_exist.marshal())
  }
  response.Obj = string(user.Marshal())
  response.Error = false // по идее не надо тк дефолт значение
  return response.Marshal()
}

var ActionUpdateUserByUser           ActionFunc = func(message Message) []byte {
  response := Response{Type: message.Type}
  var user models.User
  user, err := user.Unmarshal([]byte(message.Obj))
  if err != nil{
    response.Error = true
    response.Obj = string(incorrect_data.marshal())
    return response.Marshal()
  }
  err = userRepo.UpdateUserByUser(user) 
  if err != nil {
    response.Error = true
    response.Obj = string(not_found_err.marshal())
    return response.Marshal()
  }

  response.Error = false // по идее не надо тк дефолт значение
  return response.Marshal()
}

var ActionUpdateUserByAdmin          ActionFunc = func(message Message) []byte {
  response := Response{Type: message.Type}
  var user models.User
  user, err := user.Unmarshal([]byte(message.Obj))
  if err != nil{
    response.Error = true
    response.Obj = string(incorrect_data.marshal())
    return response.Marshal()
  }
  err = userRepo.UpdateUserByUser(user) 
  if err != nil {
    response.Error = true
    response.Obj = string(not_found_err.marshal())
    return response.Marshal()
  }

  return response.Marshal()
}

// type taskAndEmail struct {
//   Task models.Task `json:"task"`
//   Email string `json:"email"`
// }

var ActionCreateTask                 ActionFunc = func(message Message) []byte {
  var task models.Task
  response := Response{Type: message.Type}
  err := json.Unmarshal([]byte(message.Obj), &task)
  if err != nil{
    response.Error = true
    response.Obj = string(incorrect_data.marshal())
    return response.Marshal()
  }
  task, err = taskRepo.CreateTask(task) // надо чекнуть, что данные не перезатруться при создании с таким же мылом
  if err != nil {
    response.Error = true
    response.Obj = string(not_found_email.marshal())
    return response.Marshal()
  }
  
  response.Obj = string(task.Marshal())
  response.Error = false // по идее не надо тк дефолт значение
  return response.Marshal()
}

var ActionDeleteTask                 ActionFunc = func(message Message) []byte {
  var task_id uint64
  response := Response{Type: message.Type}
  err := json.Unmarshal([]byte(message.Obj), &task_id)
  if err != nil{
    response.Error = true
    response.Obj = string(incorrect_data.marshal())
    return response.Marshal()
  }
  err = taskRepo.DeleteTask(task_id) 
  if err != nil {
    response.Error = true
    response.Obj = string(not_found_email.marshal())
    return response.Marshal()
  }
  response.Obj = message.Obj
  response.Error = false // по идее не надо тк дефолт значение
  return response.Marshal()
}

var ActionUpdateTask                 ActionFunc = func(message Message) []byte {
  response := Response{Type: message.Type}
  var task models.Task
  task, err := task.Unmarshal([]byte(message.Obj))
  if err != nil{
    response.Error = true
    response.Obj = string(incorrect_data.marshal())
    return response.Marshal()
  }
  err = taskRepo.UpdateTask(task) 
  if err != nil {
    response.Error = true
    response.Obj = string(not_found_err.marshal())
    return response.Marshal()
  }

  response.Obj = message.Obj
  response.Error = false // по идее не надо тк дефолт значение
  return response.Marshal()
}
// надо проверить
var ActionFindAllTaskForUser         ActionFunc = func(message Message) []byte {
  response := Response{Type: message.Type}
  var email string
  err := json.Unmarshal([]byte(message.Obj), &email)
  if err != nil{
    response.Error = true
    response.Obj = string(incorrect_data.marshal())
    return response.Marshal()
  }
  tasks := taskRepo.FindAllActiveTaskByUserEmail(email)
  if len(tasks) == 0{
    response.Error = true
    response.Obj = string(not_found_err.marshal())
    return response.Marshal()
  }
  tasks_bytes, err := json.Marshal(tasks)
  if err != nil{
    log.Println("err in get all tasks marshal!! Err: ", err)
    response.Error = true
    response.Obj = string(not_found_err.marshal())
    return response.Marshal()
  }
  response.Obj = string(tasks_bytes)
  response.Error = false // по идее не надо тк дефолт значение
  return response.Marshal()
}
// надо проверить и получше переделать
var ActionCreateTaskForMultipleUsers ActionFunc = func(message Message) []byte {
  response := Response{Type: message.Type}
  var taskAndEmails struct{
    Task models.Task `json:"task"`
    Emails []string `json:"emails"`
  }
  err := json.Unmarshal([]byte(message.Obj), &taskAndEmails)
  if err != nil {
    log.Println("err in unmarshal emails!! Err: ", err)
    response.Error = true
    response.Obj = string(not_found_err.marshal())
    return response.Marshal()
  }
  tasksAndErrors := CreateTaskForMultipleUsers(taskAndEmails.Task, taskAndEmails.Emails)
  // if isError {
  //   response.Error = true
    // response.Obj = string(not_found_email.marshal())
    // return response.Marshal()
  // }
  // response.Error = false // по идее не надо тк дефолт значение
  bytes, err := json.Marshal(tasksAndErrors)
  if err != nil {
    // скорее всего надо будет удалить эту проверку, пока для теста
    log.Println("err in marshal tasksAndErrors: ", err)
    response.Error = true
    response.Obj = string(incorrect_data.marshal())
    return response.Marshal()
  }
  response.Obj = string(bytes)
  return response.Marshal()
}

var ActionFindAllUsers ActionFunc = func(message Message) []byte {
    response := Response{Type: message.Type}
    
    users := userRepo.FindAllActiveUsers()
    if len(users) == 0 {
      log.Println("len users = 0")
      response.Error = true
      response.Obj = string(not_found_err.marshal())
      return response.Marshal()  
    }
    bytes, err := json.Marshal(users)
    if err != nil {
      log.Printf("err in marshal users arr: %s\n", err)
      response.Error = true
      response.Obj = string(not_found_err.marshal())
      return response.Marshal()
    }
    response.Error = false
    response.Obj = string(bytes)
    return response.Marshal()
}
