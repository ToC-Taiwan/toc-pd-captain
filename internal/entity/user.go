// Package entity package entity
package entity

// User User
type User struct {
	ID       int64  `json:"-"`
	UserName string `json:"username"`
	Password string `json:"password"`
}
