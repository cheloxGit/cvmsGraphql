package resolvers

//User struct
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

//Price struct
type Price struct {
	Label    string   `json:"label"`
	Number   string   `json:"number"`
	Lapse    string   `json:"lapse"`
	Desc     []string `json:"desc"`
	LabelBuy string   `json:"labelBuy"`
}

//Service struct
type Service struct {
	TitleOne    string `json:"titleOne"`
	SubTitleOne string `json:"subTitleOne"`
	DescOne     string `json:"descOne"`
	SrcBgOne    string `json:"srcBgOne"`
	TitleTwo    string `json:"titleTwo"`
	SubTitleTwo string `json:"subTitleTwo"`
	DescTwo     string `json:"descTwo"`
	SrcBgTwo    string `json:"srcBgTwo"`
}

//FactsAboutMe struct
type FactsAboutMe struct {
	Label string `json:"label"`
	Desc  string `json:"desc"`
	Icon  string `json:icon`
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
	Rating   int    `json:"rating"`
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
	BgImage         string          `json:"bgImage"`
	MyImage         string          `json:"myImage"`
	MenuList        []string        `json:"menuList"`
	AboutMe         string          `json:"aboutMe"`
	LblAboutMe      string          `json:"lblAboutMe"`
	FactsAboutMe    []FactsAboutMe  `json:"factsAboutMe"`
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
