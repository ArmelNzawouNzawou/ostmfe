package people

type People struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	BirthDate  string `json:"birth_date"`
	Origine    string `json:"origine"`
	History    string `json:"history"`
	Profession string `json:"profession"`
}

type People_image struct {
	PeopleId  string `json:"people_id"`
	ImageId   string `json:"image_id"`
	ImageType string `json:"image_type"`
}

type People_profession struct {
	Profession  string `json:"profession"`
	People_id   string `json:"people_id"`
	Description string `json:"description"`
}

type PeoplePlace struct {
	PlaceId  string `json:"place_id"`
	PeopleId string `json:"people_id"`
}

type Profession struct {
	Id          string `json:"id"`
	Profession  string `json:"profession"`
	Description string `json:"description"`
}

type Profession_image struct {
	ProfessionId string `json:"professionId"`
	ImageId      string `json:"image_id"`
	Description  string `json:"description"`
}
