package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cheloxGit/cvmsGraphql/data"
	"github.com/cheloxGit/cvmsGraphql/schema"
	"github.com/graphql-go/graphql"
)

type reqBody struct {
	Query string `json:"query"`
}

//GqlHandler func
func GqlHandler() http.Handler {
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
		fmt.Fprintf(w, "%s", ProcessQuery(token, rBody.Query))

	})
}

//ProcessQuery func
func ProcessQuery(token string, query string) (result string) {

	retrieveCVMS := data.RetrieveCVMSFromFile()

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
