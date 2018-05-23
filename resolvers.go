package main

import (
	"context"
	"fmt"

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
