package people_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/people"
)

const professionImg = api.BASE_URL + "professionImg"

func CreateProfessionImage(prI people.Profession_image) (people.Profession_image, error) {

	entity := people.Profession_image{}
	resp, _ := api.Rest().SetBody(prI).Post(professionImg + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func UpdateProfessionImage(prI people.Profession_image) (people.Profession_image, error) {

	entity := people.Profession_image{}
	resp, _ := api.Rest().SetBody(prI).Post(professionImg + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadProfessionImage(id string) (people.Profession_image, error) {

	entity := people.Profession_image{}
	resp, _ := api.Rest().Get(professionImg + "read?id" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func DeleteProfessionImage(id string) (people.Profession_image, error) {

	entity := people.Profession_image{}
	resp, _ := api.Rest().Get(professionImg + "delete?id" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadProfessionImages() (people.Profession_image, error) {

	entity := people.Profession_image{}
	resp, _ := api.Rest().Get(professionImg + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
