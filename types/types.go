package types

import (
	"github.com/graphql-go/graphql"
)

var PriceType = graphql.NewObject(
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
			"labelBuy": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var ServiceType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Service",
		Fields: graphql.Fields{
			"titleOne": &graphql.Field{
				Type: graphql.String,
			},
			"subTitleOne": &graphql.Field{
				Type: graphql.String,
			},
			"descOne": &graphql.Field{
				Type: graphql.String,
			},
			"srcBgOne": &graphql.Field{
				Type: graphql.String,
			},
			"titleTwo": &graphql.Field{
				Type: graphql.String,
			},
			"subTitleTwo": &graphql.Field{
				Type: graphql.String,
			},
			"descTwo": &graphql.Field{
				Type: graphql.String,
			},
			"srcBgTwo": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
var CodingSkillType = graphql.NewObject(
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
var LanguageSkillType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "LanguageSkill",
		Fields: graphql.Fields{
			"language": &graphql.Field{
				Type: graphql.String,
			},
			"rating": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)
var LinearSkillType = graphql.NewObject(
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
var ExperienceType = graphql.NewObject(
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
var EducationType = graphql.NewObject(
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
var TestimonialType = graphql.NewObject(
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
var ClientType = graphql.NewObject(
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
var FunFactType = graphql.NewObject(
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
var FamType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "FactsAboutMe",
		Fields: graphql.Fields{
			"label": &graphql.Field{
				Type: graphql.String,
			},
			"desc": &graphql.Field{
				Type: graphql.String,
			},
			"icon": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var CvType = graphql.NewObject(
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
			"bgImage": &graphql.Field{
				Type: graphql.String,
			},
			"myImage": &graphql.Field{
				Type: graphql.String,
			},
			"menuList": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"aboutMe": &graphql.Field{
				Type: graphql.String,
			},
			"lblAboutMe": &graphql.Field{
				Type: graphql.String,
			},
			"factsAboutMe": &graphql.Field{
				Type: graphql.NewList(FamType),
			},
			"lblMyServices": &graphql.Field{
				Type: graphql.String,
			},
			"services": &graphql.Field{
				Type: graphql.NewList(ServiceType),
			},
			"lblPricing": &graphql.Field{
				Type: graphql.String,
			},
			"price": &graphql.Field{
				Type: graphql.NewList(PriceType),
			},
			"lblFunFacts": &graphql.Field{
				Type: graphql.String,
			},
			//TODO
			"funFacts": &graphql.Field{
				Type: graphql.NewList(FunFactType),
			},
			"lblClients": &graphql.Field{
				Type: graphql.String,
			},
			"clients": &graphql.Field{
				Type: graphql.NewList(ClientType),
			},
			"lblTestimonials": &graphql.Field{
				Type: graphql.String,
			},
			"testimonials": &graphql.Field{
				Type: graphql.NewList(TestimonialType),
			},
			"lblResume": &graphql.Field{
				Type: graphql.String,
			},
			"lblExperience": &graphql.Field{
				Type: graphql.String,
			},
			"experience": &graphql.Field{
				Type: graphql.NewList(ExperienceType),
			},
			"lblEducation": &graphql.Field{
				Type: graphql.String,
			},
			"education": &graphql.Field{
				Type: graphql.NewList(EducationType),
			},
			"lblMySkills": &graphql.Field{
				Type: graphql.String,
			},
			"lblDesign": &graphql.Field{
				Type: graphql.String,
			},
			"linearSkills": &graphql.Field{
				Type: graphql.NewList(LinearSkillType),
			},
			"lblLanguage": &graphql.Field{
				Type: graphql.String,
			},
			"languageSkills": &graphql.Field{
				Type: graphql.NewList(LanguageSkillType),
			},
			"lblCoding": &graphql.Field{
				Type: graphql.String,
			},
			"codingSkills": &graphql.Field{
				Type: graphql.NewList(CodingSkillType),
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
