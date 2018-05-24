package main

import (
	"context"
	"fmt"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/jinzhu/gorm"
)

// User is the base user model to be used throughout the app
type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
}

// DB METHODS ==========================================================================

func (db *DB) getUser(ctx context.Context, id int32) (*User, error) {
	var user User
	err := db.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	fmt.Println(user.Username)
	return &user, nil
}

func (u *User) ID(ctx context.Context) *graphql.ID {
	return gqlIDP(u.Model.ID)
}

func (u *User) USERNAME(ctx context.Context) *string {
	return &u.Username
}

func (u *User) EMAIL(ctx context.Context) *string {
	return &u.Email
}
