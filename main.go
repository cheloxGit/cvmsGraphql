package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/cheloxGit/cvmsGraphql/auth"
	"github.com/cheloxGit/cvmsGraphql/resolvers"
	"github.com/cheloxGit/cvmsGraphql/schema"
	"github.com/friendsofgo/graphiql"
	"github.com/graphql-go/graphql"
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

	http.Handle("/graphql", gqlHandler())
	http.Handle("/graphiql", graphiqlHandler)
	http.HandleFunc("/login", auth.CreateTokenEndpoint)
	http.ListenAndServe(":3000", nil)
}

func gqlHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Body == nil {
			http.Error(w, "No query data", 400)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=ascii")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
		// enableCors(&w)
		w.WriteHeader(http.StatusOK)
		var rBody reqBody
		token := r.URL.Query().Get("token")
		err := json.NewDecoder(r.Body).Decode(&rBody)
		if err != nil {
			http.Error(w, "Error parsing JSON request body", 400)
		}
		// fmt.Println(rBody.Query)
		fmt.Fprintf(w, "%s", processQuery(token, rBody.Query))

	})
}

// func enableCors(w *http.ResponseWriter) {
// 	(*w).Header().Set("Access-Control-Allow-Origin", "*")
// }

func processQuery(token string, query string) (result string) {

	retrieveCVMS := retrieveCVMSFromFile()

	params := graphql.Params{
		Schema:        schema.GqlSchema(retrieveCVMS),
		RequestString: query,
		Context:       context.WithValue(context.Background(), "token", token),
	}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		fmt.Printf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)

	return fmt.Sprintf("%s", rJSON)
}

//Open the file data.json and retrieve json data
func retrieveCVMSFromFile() func() []resolvers.CV {
	return func() []resolvers.CV {
		jsonf, err := os.Open("data.json")

		if err != nil {
			fmt.Printf("failed to open json file, error: %v", err)
		}

		jsonDataFromFile, _ := ioutil.ReadAll(jsonf)
		defer jsonf.Close()

		var cvmsData []resolvers.CV

		err = json.Unmarshal(jsonDataFromFile, &cvmsData)

		if err != nil {
			fmt.Printf("failed to parse json, error: %v", err)
		}
		return cvmsData
	}
}
