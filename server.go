package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/graph-gophers/graphql-go"

	"github.com/graph-gophers/graphql-go/relay"
)

// nolint: gas
func main() {
	s, err := getSchema("./schema.graphql")
	if err != nil {
		panic(err)
	}

	schema := graphql.MustParseSchema(s, &Resolver{})

	http.Handle("/query", &relay.Handler{Schema: schema})

	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}

func getSchema(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
