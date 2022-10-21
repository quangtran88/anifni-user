package domain

import (
	"fmt"
)

type User struct {
	Id            ID
	Pid           PID
	Name          string
	LastName      string
	FirstName     string
	Email         string
	Password      string
	Blocked       bool
	BlockedReason string
	Kyc           KYC
	Investment    UserInvestment
	Staff         Staff
}

func NewUser(in CreateUserInput) *User {
	return &User{
		Pid:       in.Pid,
		Name:      fmt.Sprintf("%s %s", in.FirstName, in.LastName),
		LastName:  in.LastName,
		FirstName: in.FirstName,
		Email:     in.Email,
		Password:  in.Password,
		Blocked:   false,
	}
}

type CreateUserInput struct {
	Pid       PID
	Email     string
	Password  string
	LastName  string
	FirstName string
}
