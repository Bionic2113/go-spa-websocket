package models

import (
  "encoding/json"
)

type Task struct {
	Task_id     uint64 `json:"task_id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"` // может лучше хранить числом 🤔
	Time_start  int64  `json:"time_start"`
	Time_finish int64  `json:"time_finish"`
  User_email  string `json:"user_email,omitempty"`
}

// а нужны ли они тут, ведь клиент будет отправлять?
const ( // набор статусов
	Pending   = "pending"   // ожидает
	Overdue   = "overdue"   // просрочено
	Completed = "completed" // выполнено
	Deleted   = "deleted"   // удалено
)

func (t Task) Marshal() []byte{
  bytes, _ := json.Marshal(t)
  return bytes
}

func (t Task) Unmarshal(bytes []byte) (Task, error) {
  err := json.Unmarshal(bytes, &t)
  return t, err
}
