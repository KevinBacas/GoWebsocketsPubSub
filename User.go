package main

import (
  "golang.org/x/net/websocket"
  "github.com/KevinBacas/go-uuid/uuid"
  log "github.com/Sirupsen/logrus"
)

type User struct {
  UId string // User identifier
  ws* websocket.Conn // WebSocket used to communicate with him
}

func NewUser(webs* websocket.Conn) *User {
  log.Info("Creating User")
  uuid := uuid.New()
  return &User{
    ws: webs,
    UId: uuid,
  }
}

func (this *User) GetUId() string {
  return this.UId
}

func (this *User) SendMessage(message *MessageResponse) error {
  log.Info("User:", "Sending message to :", this.UId)
  err := websocket.JSON.Send(this.ws, message)
  if err != nil {
    log.Error("User:", "an error occured:", err)
  }
  return err
}
