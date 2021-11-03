package domain

import (
	"errors"
	"fmt"
)

type User struct {
    name UserName
    phoneNumber UserPhoneNumber
}

type UserName struct {
    name string
}

type UserPhoneNumber struct {
    phoneNumber string
}

func NewUser(name string, phoneNumber string) (*User, error) {
    userName, err := NewUserName(name)
    if err != nil {
	return nil, err
    }

    return &User {
	name: *userName,
	phoneNumber: NewUserPhoneNumber(phoneNumber),
    }, nil
}

func NewUserName(name string) (*UserName, error) {
    if len(name) < 3 {
	return nil, errors.New("user name should be at least 3 character")
    }

    return &UserName{
	name: name,
    }, nil
}

func NewUserPhoneNumber(phoneNumber string) UserPhoneNumber {
    return UserPhoneNumber{
	phoneNumber: phoneNumber,
    }
}

func (u User) String() string {
   return fmt.Sprintf("User [name=%s, phoneNumber=%s]", u.name, u.phoneNumber)
}
