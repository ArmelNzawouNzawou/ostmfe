package member

type Member struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Birthday  string `json:"birthday"`
	ImageCode string `json:"image_code"`
}
