package main

import (
  "bytes"

  log "github.com/Sirupsen/logrus"
  "github.com/valyala/fasthttp"
)

var baseURL string  = "/go/src/hello/"
var conn *ConnectionManager

var (
  staticPrefix  = []byte("/static/")
  staticHandler = fasthttp.FSHandler(baseURL+"static/", 1)
)

// func EchoServer(ws *websocket.Conn) {
// 	log.WithFields(log.Fields{
//     "object": "Server",
//   }).Info("Connected!")
//   log.WithFields(log.Fields{
//     "object": "Server",
//   }).Info("Adding user to ConnectionManager")
//   err := conn.AddUser(NewUser(ws))
//   if(err == nil) {
//     var err error = nil
//     for err == nil {
//       var message MessageRequest
//       err = websocket.JSON.Receive(ws, &message)
//       if err == nil {
//         log.WithFields(log.Fields{
//           "object": "Server",
//           "message": message,
//         }).Info("Received from client")
//         reply := NewMessageResponse(message.Message)
//         log.WithFields(log.Fields{
//           "object": "Server",
//           "reply": reply,
//         }).Info("Sending to clients")
//         conn.Broadcast(reply)
//       } else {
//         log.WithFields(log.Fields{
//           "object": "Server",
//           "error": err,
//         }).Error("Can't receive")
//       }
//     }
//   } else {
//     log.Error(err)
//   }
// }

// Main request handler
func requestHandler(ctx *fasthttp.RequestCtx) {
    path := ctx.Path()
    switch {
    case bytes.HasPrefix(path, staticPrefix):
        staticHandler(ctx)
    default:
        staticHandler(ctx)
    }
}

func main() {
  // Configuring Logger
  log.SetFormatter(&log.TextFormatter{
    ForceColors: true,
  })
  // log.SetLevel(log.WarnLevel)

  log.Info("Starting server...")
  err := fasthttp.ListenAndServe(":8000", requestHandler)
  if err != nil {
    log.Fatalf("Error in server: %s", err)
  }
}
