package transport

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"sushkof/internal/services"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool { // нужно для работы в бразуере по ws
		return true
	},
}

func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Ошибка обновления соединения:", err)
		return
	}
	defer conn.Close()

	var message services.Message
	// for надо делать
	for {
		err = conn.ReadJSON(&message)
		if err != nil {
			log.Printf("корявый json, err: %s\n", err)
      break
			//continue - а лучше возвращай ошибку для отображения
		}
		log.Println("message type: ", message.Type)
		action, ok := services.Actions[message.Type]
		if !ok {
			log.Println("нет в мапе такой кодовой фразы")
			//continue - а лучше возвращай ошибку для отображения
		}
		bytes := action.MakeAction(message) // надо проверить, может и без мейка надо

		err = conn.WriteMessage(websocket.TextMessage, bytes) //WriteJSON(bytes) // надо проверить, если что то по обычному отправлять
		if err != nil {
			log.Println("err in write json: ", err)
      continue
		}
		log.Println("response send")
	}
}
