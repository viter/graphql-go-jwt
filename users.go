package main

import (
	"context"

	"github.com/jinzhu/gorm"
)

// User is the base user model to be used throughout the app
type User struct {
	gorm.Model
	Username string
}

// DB METHODS ==========================================================================

func (db *DB) getUser(ctx context.Context, id int32) (*User, error) {
	var user User
	err := db.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
