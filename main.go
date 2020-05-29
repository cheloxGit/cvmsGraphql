package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/friendsofgo/graphiql"
	"github.com/graphql-go/graphql"
	"github.com/mitchellh/mapstructure"
)

//User struct
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

//Price struct
type Price struct {
	Label  string   `json:"label"`
	Number string   `json:"number"`
	Lapse  string   `json:"lapse"`
	Desc   []string `json:"descBasic"`
}

//Service struct
type Service struct {
	Title    string `json:"title"`
	SubTitle string `json:"subTitle"`
	Desc     string `json:"desc"`
	SrcBg    string `json:"srcBg"`
}

//FactsAboutMe struct
type FactsAboutMe struct {
	LblAge       string `json:"lblAge"`
	Age          string `json:"age"`
	LblResidence string `json:"lblResidence"`
	Residence    string `json:"residence"`
	LblState     string `json:"lblState"`
	State        string `json:"state"`
}

//CodingSkill struct
type CodingSkill struct {
	NumberOne string `json:"numberOne"`
	LabelOne  string `json:"labelOne"`
	NumberTwo string `json:"numberTwo"`
	LabelTwo  string `json:"labelTwo"`
}

//LanguageSkill struct
type LanguageSkill struct {
	Language string `json:"language"`
	Rating   string `json:"rating"`
}

//LinearSkill struct
type LinearSkill struct {
	Name    string `json:"name"`
	Percent string `json:"percent"`
}

//Experience struct
type Experience struct {
	Lapse    string `json:"lapse"`
	Position string `json:"position"`
	Title    string `json:"title"`
	Desc     string `json:"desc"`
}

//Education struct
type Education struct {
	Lapse    string `json:"lapse"`
	Position string `json:"position"`
	Title    string `json:"title"`
	Desc     string `json:"desc"`
}

//Testimonial struct
type Testimonial struct {
	Name        string `json:"name"`
	Testimonial string `json:"testimonial"`
}

//Client struct
type Client struct {
	Client string `json:"client"`
	SrcBg  string `json:"srcBg"`
}

//FunFact struct
type FunFact struct {
	Number string   `json:"number"`
	Desc   []string `json:"desc"`
}

//CV struct
type CV struct {
	ID              int             `json:"id"`
	FullName        string          `json:"fullName"`
	Degree          string          `json:"degree"`
	MenuList        []string        `json:"menuList"`
	AboutMe         string          `json:"aboutMe"`
	FactsAboutMe    FactsAboutMe    `json:"factsAboutMe"`
	LblMyServices   string          `json:"lblMyServices"`
	Services        []Service       `json:"services"`
	LblPricing      string          `json:"lblPricing"`
	Price           []Price         `json:"price"`
	LblFunFacts     string          `json:"lblFunFacts"`
	FunFacts        []FunFact       `json:"FunFacts"`
	LblClients      string          `json:"lblClients"`
	Clients         []Client        `json:"clients"`
	LblTestimonials string          `json:"lblTestimonials"`
	Testimonials    []Testimonial   `json:"testimonials"`
	LblResume       string          `json:"lblResume"`
	LblExperience   string          `json:"lblExperience"`
	Experience      []Experience    `json:"experience"`
	LblEducation    string          `json:"lblEducation"`
	Education       []Education     `json:"education"`
	LblMySkills     string          `json:"lblMySkills"`
	LblDesign       string          `json:"lblDesign"`
	LinearSkills    []LinearSkill   `json:"linearSkills"`
	LblLanguage     string          `json:"lblLanguage"`
	LanguageSkills  []LanguageSkill `json:"languageSkills"`
	LblCoding       string          `json:"lblCoding"`
	CodingSkills    []CodingSkill   `json:"codingSkills"`
	KnowledgeSkills []string        `json:"knowledgeSkills"`
	LblKnowledge    string          `json:"lblKnowledge"`
}

var priceType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Price",
		Fields: graphql.Fields{
			"label": &graphql.Field{
				Type: graphql.String,
			},
			"number": &graphql.Field{
				Type: graphql.String,
			},
			"lapse": &graphql.Field{
				Type: graphql.String,
			},
			"desc": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
		},
	},
)

var serviceType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Service",
		Fields: graphql.Fields{
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"subTitle": &graphql.Field{
				Type: graphql.String,
			},
			"desc": &graphql.Field{
				Type: graphql.String,
			},
			"srcBg": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
var codingSkillType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "CodingSkill",
		Fields: graphql.Fields{
			"numberOne": &graphql.Field{
				Type: graphql.String,
			},
			"labelOne": &graphql.Field{
				Type: graphql.String,
			},
			"numberTwo": &graphql.Field{
				Type: graphql.String,
			},
			"labelTwo": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
var languageSkillType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "LanguageSkill",
		Fields: graphql.Fields{
			"language": &graphql.Field{
				Type: graphql.String,
			},
			"rating": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
var linearSkillType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "LinearSkill",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"percent": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
var experienceType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Experience",
		Fields: graphql.Fields{
			"lapse": &graphql.Field{
				Type: graphql.String,
			},
			"position": &graphql.Field{
				Type: graphql.String,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"desc": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
var educationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Education",
		Fields: graphql.Fields{
			"lapse": &graphql.Field{
				Type: graphql.String,
			},
			"position": &graphql.Field{
				Type: graphql.String,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"desc": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
var testimonialType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Testimonial",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"testimonial": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
var clientType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Client",
		Fields: graphql.Fields{
			"client": &graphql.Field{
				Type: graphql.String,
			},
			"srcBg": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
var funFactType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "FunFact",
		Fields: graphql.Fields{
			"number": &graphql.Field{
				Type: graphql.String,
			},
			"desc": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
		},
	},
)
var famType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "FactsAboutMe",
		Fields: graphql.Fields{
			"lblAge": &graphql.Field{
				Type: graphql.String,
			},
			"age": &graphql.Field{
				Type: graphql.String,
			},
			"lblResidence": &graphql.Field{
				Type: graphql.String,
			},
			"residence": &graphql.Field{
				Type: graphql.String,
			},
			"lblState": &graphql.Field{
				Type: graphql.String,
			},
			"state": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var cvType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "CV",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"fullName": &graphql.Field{
				Type: graphql.String,
			},
			"degree": &graphql.Field{
				Type: graphql.String,
			},
			"menuList": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"aboutMe": &graphql.Field{
				Type: graphql.String,
			},
			"factsAboutMe": &graphql.Field{
				Type: famType,
			},
			"lblMyServices": &graphql.Field{
				Type: graphql.String,
			},
			"services": &graphql.Field{
				Type: graphql.NewList(serviceType),
			},
			"lblPricing": &graphql.Field{
				Type: graphql.String,
			},
			"price": &graphql.Field{
				Type: graphql.NewList(priceType),
			},
			"lblFunFacts": &graphql.Field{
				Type: graphql.String,
			},
			//TODO
			"funFacts": &graphql.Field{
				Type: graphql.NewList(funFactType),
			},
			"lblClients": &graphql.Field{
				Type: graphql.String,
			},
			"clients": &graphql.Field{
				Type: graphql.NewList(clientType),
			},
			"lblTestimonials": &graphql.Field{
				Type: graphql.String,
			},
			"testimonials": &graphql.Field{
				Type: graphql.NewList(testimonialType),
			},
			"lblResume": &graphql.Field{
				Type: graphql.String,
			},
			"lblExperience": &graphql.Field{
				Type: graphql.String,
			},
			"experience": &graphql.Field{
				Type: graphql.NewList(experienceType),
			},
			"lblEducation": &graphql.Field{
				Type: graphql.String,
			},
			"education": &graphql.Field{
				Type: graphql.NewList(educationType),
			},
			"lblMySkills": &graphql.Field{
				Type: graphql.String,
			},
			"lblDesign": &graphql.Field{
				Type: graphql.String,
			},
			"linearSkills": &graphql.Field{
				Type: graphql.NewList(linearSkillType),
			},
			"lblLanguage": &graphql.Field{
				Type: graphql.String,
			},
			"languageSkills": &graphql.Field{
				Type: graphql.NewList(languageSkillType),
			},
			"lblCoding": &graphql.Field{
				Type: graphql.String,
			},
			"codingSkills": &graphql.Field{
				Type: graphql.NewList(codingSkillType),
			},
			"knowledgeSkills": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"lblKnowledge": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var jwtSecret []byte = []byte("thepolyglotdeveloper")

type reqBody struct {
	Query string `json:"query"`
}

//ValidateJWT func
func ValidateJWT(t string) (interface{}, error) {
	if t == "" {
		return nil, errors.New("Authorization token must be present")
	}
	token, _ := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return jwtSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var decodedToken interface{}
		mapstructure.Decode(claims, &decodedToken)
		return decodedToken, nil
	} else {
		return nil, errors.New("Invalid authorization token")
	}
}

//CreateTokenEndpoint func
func CreateTokenEndpoint(response http.ResponseWriter, request *http.Request) {
	var user User
	_ = json.NewDecoder(request.Body).Decode(&user)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"password": user.Password,
	})
	tokenString, error := token.SignedString(jwtSecret)
	if error != nil {
		fmt.Println(error)
	}
	response.Header().Set("content-type", "application/json")
	response.Write([]byte(`{ "token": "` + tokenString + `" }`))
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
	http.HandleFunc("/login", CreateTokenEndpoint)
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
func retrieveCVMSFromFile() func() []CV {
	return func() []CV {
		jsonf, err := os.Open("data.json")

		if err != nil {
			fmt.Printf("failed to open json file, error: %v", err)
		}

		jsonDataFromFile, _ := ioutil.ReadAll(jsonf)
		defer jsonf.Close()

		var cvmsData []CV

		err = json.Unmarshal(jsonDataFromFile, &cvmsData)

		if err != nil {
			fmt.Printf("failed to parse json, error: %v", err)
		}
		return cvmsData
	}
}

// Define the GraphQL Schema
// func gqlSchema(queryCV func() []CV) graphql.Schema {
// 	fields := graphql.Fields{
// 		"cvms": &graphql.Field{
// 			Type:        graphql.NewList(cvType),
// 			Description: "All CVs",
// 			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
// 				// fmt.Println("ValidateJWT: params.Context.Value(token).(string)")
// 				// fmt.Println(params.Context.Value("token").(string))
// 				// _, err := ValidateJWT(params.Context.Value("token").(string))
// 				// fmt.Println("ValidateJWT: err")
// 				// fmt.Println(err)
// 				// if err != nil {
// 				// 	return nil, err
// 				// }

// 				// fmt.Printf(err)
// 				// if err != nil {
// 				// 	fmt.Printf("CVMS Failed: %v", err)
// 				// 	return nil, err
// 				// }
// 				// for _, accountMock := range accountsMock {
// 				// 	if accountMock.Username == account.(User).Username {
// 				// 		return accountMock, nil
// 				// 	}
// 				// }
// 				fmt.Println(queryCV())
// 				return queryCV(), nil
// 			},
// 		},
// 		"cvmsid": &graphql.Field{
// 			Type:        cvType,
// 			Description: "Get CVs by ID",
// 			Args: graphql.FieldConfigArgument{
// 				"id": &graphql.ArgumentConfig{
// 					Type: graphql.Int,
// 				},
// 			},
// 			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
// 				id, success := params.Args["id"].(int)
// 				if success {
// 					for _, cv := range queryCV() {
// 						if int(cv.ID) == id {
// 							return cv, nil
// 						}
// 					}
// 				}
// 				return nil, nil
// 			},
// 		},
// 	}
// 	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
// 	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
// 	schema, err := graphql.NewSchema(schemaConfig)
// 	if err != nil {
// 		fmt.Printf("failed to create new schema, error: %v", err)
// 	}

// 	return schema

// }

// {"query":"{\n  cvms{\n    lblCoding\n  }\n}","variables":null,"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXNzd29yZCI6IiIsInVzZXJuYW1lIjoiIn0.b1gvcw1zQS2mCz3cwWbx6qU_syYJI7ZRwIfgKi9hcXI"}
// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXNzd29yZCI6InNvLXNlY3JldCIsInVzZXJuYW1lIjoiY2hlbG94In0.SEJ6tIM89KEdPNxUfOT9RuONG9GZJ3dFb5Y1bjcpraA
