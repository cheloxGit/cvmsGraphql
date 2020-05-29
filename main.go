package main

import (
	"fmt"
	"net/http"

	"github.com/cheloxGit/cvmsGraphql/auth"
	"github.com/cheloxGit/cvmsGraphql/handler"
	"github.com/friendsofgo/graphiql"
)

type reqBody struct {
	Query string `json:"query"`
}

func main() {
	// c := cors.New(cors.Options{
	// 	AllowedOrigins: []string{"http://foo.com"},
	// })
	fmt.Println("Starting the application at :3000...")
	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/graphql")
	if err != nil {
		panic(err)
	}

	http.Handle("/graphql", handler.GqlHandler())
	http.Handle("/graphiql", graphiqlHandler)
	http.HandleFunc("/login", auth.CreateTokenEndpoint)
	http.ListenAndServe(":3000", nil)
}
