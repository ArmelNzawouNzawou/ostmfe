package project

type Project struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageCode   string `json:"image_code"`
	Profile     string `json:"profile"`
}

type ProjectMember struct {
	ProjectId   string `json:"project_id"`
	MemberId    string `json:"member_id"`
	Description string `json:"description"`
}

type ProjectPartner struct {
	ProjectId   string `json:"projectId"`
	PartenerID  string `json:"partenerId"`
	Description string `json:"description"`
}
