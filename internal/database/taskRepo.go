package database

import (
	"database/sql"
	"log"
	"sushkof/internal/models"
)

type TaskRepo interface {
	FindAllActiveTaskByUserEmail(email string) []models.Task
	CreateTask(task models.Task) (models.Task, error)
	DeleteTask(task_id uint64) error
	UpdateTask(task models.Task) error
}

type TaskRepoDB struct {
	*sql.DB
}

func (tdb TaskRepoDB) FindAllActiveTaskByUserEmail(email string) []models.Task {
	var tasks []models.Task
	rows, err := tdb.DB.Query("select task_id, title, description, status, time_start, time_finish, user_email from tasks where status != 'deleted' and user_email = ? order by time_start", email)
	if err != nil {
		log.Println("err in query: ", err)
		return tasks
	}

	for rows.Next() {
		task := models.Task{}
		err := rows.Scan(&task.Task_id, &task.Title, &task.Description, &task.Status, &task.Time_start, &task.Time_finish, &task.User_email)
		if err != nil {
			log.Println("err in scan: ", err)
			return tasks
		}
		tasks = append(tasks, task)
	}
	if err = rows.Err(); err != nil {
		log.Println("err in Err: ", err)
		return tasks
	}
	return tasks
}

func (tdb TaskRepoDB) CreateTask(task models.Task) (models.Task, error) {
    result, err := tdb.DB.Exec("insert into tasks(title, description, status, time_start, time_finish, user_email) values(?,?,?,?,?,?)",
        task.Title, task.Description, task.Status, task.Time_start, task.Time_finish, task.User_email)
    if err != nil {
        return models.Task{}, err
    }
    id, err := result.LastInsertId()
    if err != nil {
        return models.Task{}, err
    }
    task.Task_id = uint64(id)
    return task, nil
}


func (tdb TaskRepoDB) DeleteTask(task_id uint64) error {
	_, err := tdb.Exec("update tasks set status = 'deleted' where task_id = ?", task_id)
	return err
}

func (tdb TaskRepoDB) UpdateTask(task models.Task) error {
	_, err := tdb.Exec("update tasks set title = ?, description = ?, status = ?, time_finish= ? where task_id = ?",
		task.Title, task.Description, task.Status, task.Time_finish, task.Task_id)
	return err
}
