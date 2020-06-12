package schema

import (
	"fmt"

	"github.com/cheloxGit/cvmsGraphql/auth"
	"github.com/cheloxGit/cvmsGraphql/resolvers"
	"github.com/cheloxGit/cvmsGraphql/types"
	"github.com/graphql-go/graphql"
)

// GqlSchema Define the GraphQL Schema
func GqlSchema(queryCV func() []resolvers.CV) graphql.Schema {
	fields := graphql.Fields{
		"cvms": &graphql.Field{
			Type:        graphql.NewList(types.CvType),
			Description: "All CVs",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				fmt.Println("ValidateJWT: params.Context.Value(token).(string)")
				fmt.Println(params.Context.Value("token").(string))
				_, err := auth.ValidateJWT(params.Context.Value("token").(string))
				fmt.Println("ValidateJWT: err")
				fmt.Println(err)
				if err != nil {
					return nil, err
				}

				// fmt.Printf(err)
				// if err != nil {
				// 	fmt.Printf("CVMS Failed: %v", err)
				// 	return nil, err
				// }
				// for _, accountMock := range accountsMock {
				// 	if accountMock.Username == account.(User).Username {
				// 		return accountMock, nil
				// 	}
				// }
				// fmt.Println(queryCV())
				return queryCV(), nil
			},
		},
		"cvmsid": &graphql.Field{
			Type:        types.CvType,
			Description: "Get CVs by ID",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id, success := params.Args["id"].(int)
				if success {
					for _, cv := range queryCV() {
						if int(cv.ID) == id {
							return cv, nil
						}
					}
				}
				return nil, nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		fmt.Printf("failed to create new schema, error: %v", err)
	}

	return schema

}
