package transport

import(
  "net/http"
  "log"
  "fmt"
)

func StartServer(port int) {
  http.HandleFunc("/main.js", jsHandler)
  http.HandleFunc("/main_new.js", jsTestHandler)
  http.HandleFunc("/auth", authHandler)
  http.HandleFunc("/register", registerHandler)
  http.HandleFunc("/login", loginHandler)
  http.HandleFunc("/auth.js", authJsHandler)
  http.HandleFunc("/test.css", cssHandler)
  http.Handle("/", authMiddleware( http.HandlerFunc(startHandler)))
  http.Handle("/ws", authMiddleware( http.HandlerFunc(WebsocketHandler)))

  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}


func startHandler(w http.ResponseWriter, r *http.Request){
  http.ServeFile(w, r, "../../static/spa.html")
}

func jsHandler(w http.ResponseWriter, r *http.Request){
  http.ServeFile(w, r, "../../static/main.js")
}

func authJsHandler(w http.ResponseWriter, r *http.Request){
  http.ServeFile(w, r, "../../static/auth.js")
}

func cssHandler(w http.ResponseWriter, r *http.Request){
  http.ServeFile(w, r, "../../static/test.css")
}


//for test
func jsTestHandler(w http.ResponseWriter, r *http.Request){
  http.ServeFile(w, r, "../../static/main_new.js")
}

