package main

import (
    "github.com/stretchr/testify/mock"
    "testing"
)

type Database interface {
    GetUserName(userID int) string
}

type UserService struct {
    DB Database
}

func (s *UserService) GetUserName(userID int) string {
    return s.DB.GetUserName(userID)
}
