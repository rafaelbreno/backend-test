package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name           string
	Email          string
	Password       string
	Biography      string
	ProfilePicture string
	City           string
	State          string
}

type CreateUserInput struct {
	Name                 string `json:"name" binding:"required"`
	Email                string `json:"email" binding:"required"`
	Password             string `json:"password" binding:"required"`
	PasswordConfirmation string `json:"password_confirmation" binding:"required"`
	Biography            string `json:"biography" binding:"required"`
	ProfilePicture       string `json:"profile_picture" binding:"required"`
	City                 string `json:"city" binding:"required"`
	State                string `json:"state" binding:"required"`
}

type UpdateUserInput struct {
}
