package main

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model

	Id                 int
	Username           string `json:"username"`
	Password           string `json:"password"`
	Email              string `json:"email"`
	Enabled            bool   `json:"enabled"`
	AccountExpired     bool   `json:"account_expired"`
	CredentialsExpired bool   `json:"credentials_expired"`
	AccountNonLocked   bool   `json:"account_non_locked"`
}
