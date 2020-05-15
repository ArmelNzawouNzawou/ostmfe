package collection

type Collection struct {
	Id                 string `json:"id"`
	Name               string `json:"name"`
	ImageCode          string `json:"image_code"`
	ProfileDescription string `json:"profile_description"`
	History            string `json:"history"`
}

type Collection_image struct {
	ImageId           string `json:"image_id"`
	CollectionImageId string `json:"collection_image_id"`
	Description       string `json:"description"`
}
