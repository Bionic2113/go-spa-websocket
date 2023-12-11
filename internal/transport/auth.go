package transport

import (
	"encoding/json"
	"log"
	"net/http"
	"sushkof/internal/models"
	"sushkof/internal/services"
	"time"

	"github.com/google/uuid"
)

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("uuid_cookie")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		_, ok := services.CookiesMap[cookie.Value]
		if !ok {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		next.ServeHTTP(w, r)
	})
}

type logpass struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../../static/auth.html")
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var lg logpass
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&lg)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Некорректные данные"))
		return
	}

	user, err := services.Authentication(lg.Email, lg.Password)
	if err != nil {
		log.Println("не нашел")
		w.WriteHeader(http.StatusNotFound)
    w.Write([]byte(err.Error()))
		return
	} else {
		// надо в мапу то пользователя пихать, придется возвращать его ещё
		// в мапу запихал, но данные пользователя надо будет обновлять
		user_id := uuid.NewString()
		services.CookiesMap[user_id] = user
		http.SetCookie(w, &http.Cookie{Name: "uuid_cookie",
			Value:    user_id,
			MaxAge:   30 * 24 * 3600,
			Expires:  time.Now().Add(30 * 24 * time.Hour),
			HttpOnly: false,
			Secure:   false,
		})

		w.WriteHeader(http.StatusOK)
		}
}

type base_user struct {
	Age      uint8  `json:"age"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var bu base_user
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&bu)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
    log.Println("error in decode bu: ", err)
		w.Write([]byte("Некорректные данные"))
		return
	}

	
	w.Header().Set("Content-Type", "application/json")
	user := models.User{
		Admin:    false,
		Age:      bu.Age,
		Name:     bu.Name,
		Email:    bu.Email,
		Password: bu.Password,
	}
	bytes := services.Actions["create user"].MakeAction(
		services.Message{
			Type: "",
			Obj:  string(user.Marshal()),
		})
	user_id := uuid.NewString()
	services.CookiesMap[user_id] = &user
	http.SetCookie(w, &http.Cookie{Name: "uuid_cookie",
		Value:    user_id,
		MaxAge:   30 * 24 * 3600,
		Expires:  time.Now().Add(30 * 24 * time.Hour),
		HttpOnly: false,
		Secure:   false,
	})

	w.WriteHeader(http.StatusOK)
	w.Write(bytes) // need delete

}
