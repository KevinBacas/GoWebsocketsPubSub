package main

import (
  "time"

  "github.com/KevinBacas/go-uuid/uuid"

)

type MessageResponse struct {
  UId string
  Message string
  Timestamp time.Time
}

func NewMessageResponse(message string) *MessageResponse {
  return &MessageResponse{
    UId: uuid.New(),
    Message: message,
    Timestamp: time.Now(),
  }
}
