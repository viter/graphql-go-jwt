package main

import (
	"context"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	graphql "github.com/graph-gophers/graphql-go"
)

var db DB

func init() {
	var err error
	d, err := newDB()
	if err != nil {
		panic(err)
	}

	db = *d
}

// Resolver is the root resolver
type Resolver struct{}

// GetUser resolves the getUser query
func (r *Resolver) GetUser(ctx context.Context, args struct{ ID int32 }) (*User, error) {
	return db.getUser(ctx, int32(args.ID))
}

func gqlIDP(id uint) *graphql.ID {
	r := graphql.ID(fmt.Sprint(id))
	return &r
}

func (r *Resolver) Login(ctx context.Context, args struct {
	Email    string
	Password string
}) (*string, error) {
	var user User
	var tokenString string
	err := db.DB.Debug().Where("email like ?", args.Email).First(&user).Error
	if err != nil {
		es := ""
		return &es, err
	}

	if user.Password == args.Password {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":    user.Model.ID,
			"email": user.Email,
		})
		tokenString, error := token.SignedString([]byte(jwtSecret))
		if error != nil {
			fmt.Println(error)
		}
		return &tokenString, nil
	}
	return &tokenString, nil
}
