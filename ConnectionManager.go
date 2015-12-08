package main

import (
  "errors"

  log "github.com/Sirupsen/logrus"
)

var maxConnections int = 10000

type ConnectionManager struct {
  userList []*User
  length int
}

func NewConnectionManager(users []*User) *ConnectionManager {
  log.Info("Creating ConnectionManager")
  return &ConnectionManager{
    userList: users,
    length: len(users),
  }
}

func (this *ConnectionManager) AddUser(user *User) error {
  log.Info("ConnectionManager: Adding user")
  var res error = nil
  if(this.length < maxConnections) {
    this.userList = append(this.userList, user)
    this.length++
  } else {
    res = errors.New("Cannot add User")
  }
  return res
}

func pos(slice []*User, user *User) int {
    for p, v := range slice {
        if (v == user) {
            return p
        }
    }
    return -1
}

func (this *ConnectionManager) Broadcast(message *MessageResponse) {
  for _, user := range this.userList {
    err := user.SendMessage(message)
    if(err != nil) {
      log.Error("ConnectionManager:", "An error occured while trying to send to:", user.GetUId(), ", removing him...")
      pos := pos(this.userList, user)
      this.userList = append(this.userList[:pos], this.userList[pos+1:]...)
    }
  }
}
