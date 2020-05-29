package schema

import (
	"fmt"

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
				// fmt.Println("ValidateJWT: params.Context.Value(token).(string)")
				// fmt.Println(params.Context.Value("token").(string))
				// _, err := auth.ValidateJWT(params.Context.Value("token").(string))
				// fmt.Println("ValidateJWT: err")
				// fmt.Println(err)
				// if err != nil {
				// 	return nil, err
				// }

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
				fmt.Println(queryCV())
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

// var (
// 	Schema graphql.Schema
// )

// func init() {
// 	Query := graphql.NewObject(graphql.ObjectConfig{
// 		Name: "Query",
// 		Fields: graphql.Fields{
// 			"human": &graphql.Field{
// 				Type: types.HumanType,
// 				Args: graphql.FieldConfigArgument{
// 					"id": &graphql.ArgumentConfig{
// 						Description: "id of the human",
// 						Type:        graphql.NewNonNull(graphql.String),
// 					},
// 				},
// 				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 					// validate token
// 					isValid, err := auth.ValidateToken(p.Context.Value("token").(string))
// 					if err != nil {
// 						return nil, err
// 					}
// 					if !isValid {
// 						return nil, gqlerrors.FormatError(errors.New("Invalid token"))
// 					}

// 					// get person id
// 					id, err := strconv.Atoi(p.Args["id"].(string))
// 					if err != nil {
// 						return nil, err
// 					}

// 					char, err := resolvers.GetHuman(id)
// 					if err != nil {
// 						return nil, err
// 					}

// 					return char, nil
// 				},
// 			},
// 		},
// 	})

// 	Mutation := graphql.NewObject(graphql.ObjectConfig{
// 		Name: "Mutation",
// 		Fields: graphql.Fields{
// 			"createToken": &graphql.Field{
// 				Type:        graphql.String,
// 				Description: "creates a new  JWT token ",
// 				Args: graphql.FieldConfigArgument{
// 					"username": &graphql.ArgumentConfig{
// 						Description: "username",
// 						Type:        graphql.NewNonNull(graphql.String),
// 					},
// 					"password": &graphql.ArgumentConfig{
// 						Description: "password",
// 						Type:        graphql.NewNonNull(graphql.String),
// 					},
// 				},
// 				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
// 					// marshall and cast the argument value
// 					username, _ := params.Args["username"].(string)
// 					password, _ := params.Args["password"].(string)

// 					token, err := auth.CreateToken(username, password)
// 					if err != nil {
// 						return nil, err
// 					}

// 					return token, nil

// 				},
// 			},
// 		},
// 	})

// 	schema, err := graphql.NewSchema(graphql.SchemaConfig{
// 		Query:    Query,
// 		Mutation: Mutation,
// 	})
// 	if err != nil {
// 		panic(err)
// 	}
// 	Schema = schema
// }
