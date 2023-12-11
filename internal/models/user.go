package models

import (
  "encoding/json"
)

type User struct {
	// Active   bool   // оптимизация памяти, поэтому булеаны выше
	Admin    bool   `json:"admin"`
	Age      uint8  `json:"age"`
	Name     string `json:"name"`
	Email    string `json:"email"` // user_id не вижу смысл создавать, тк мыло будет уникальным
	Password string `json:"password"`
}

func (u User) Marshal() []byte{
  bytes, _ := json.Marshal(u)
  return bytes
}

func (u User) Unmarshal(bytes []byte) (User, error) {
  err := json.Unmarshal(bytes, &u)
  return u, err
}
