package main

import (
  "net/http"

  log "github.com/Sirupsen/logrus"
  "golang.org/x/net/websocket"
  "github.com/carbocation/interpose"
  "github.com/carbocation/interpose/middleware"
  "github.com/gorilla/mux"
)

var baseURL string  = "/go/src/hello/"
var conn *ConnectionManager

func EchoServer(ws *websocket.Conn) {
	log.WithFields(log.Fields{
    "object": "Server",
  }).Info("Connected!")
  log.WithFields(log.Fields{
    "object": "Server",
  }).Info("Adding user to ConnectionManager")
  err := conn.AddUser(NewUser(ws))
  if(err == nil) {
    var err error = nil
    for err == nil {
      var message MessageRequest
      err = websocket.JSON.Receive(ws, &message)
      if err == nil {
        log.WithFields(log.Fields{
          "object": "Server",
          "message": message,
        }).Info("Received from client")
        reply := NewMessageResponse(message.Message)
        log.WithFields(log.Fields{
          "object": "Server",
          "reply": reply,
        }).Info("Sending to clients")
        conn.Broadcast(reply)
      } else {
        log.WithFields(log.Fields{
          "object": "Server",
          "error": err,
        }).Error("Can't receive")
      }
    }
  } else {
    log.Error(err)
  }
}

func main() {
  // Configuring Logger
  log.SetFormatter(&log.TextFormatter{
    ForceColors: true,
  })
  // log.SetLevel(log.WarnLevel)

  conn = NewConnectionManager(nil)

  log.Info("Configuring main Router")
	router := mux.NewRouter()
  router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, baseURL+"index.html")
  })
  router.Handle("/echo", websocket.Handler(EchoServer))
  router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(baseURL+"static/"))))

  log.Info("Adding logging middleware")
  middle := interpose.New()
	middle.Use(middleware.GorillaLog())
	middle.UseHandler(router)

  log.Info("Starting server...")
  err := http.ListenAndServe(":8000", middle)
  if err != nil {
      panic("ListenAndServe: " + err.Error())
  }
}
