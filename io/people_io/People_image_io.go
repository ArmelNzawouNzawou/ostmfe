package people_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/people"
)

const peopleImg = api.BASE_URL + "peopleImg"

func CreatePeopleImage(pI people.People_image) (people.People_image, error) {

	entity := people.People_image{}
	resp, _ := api.Rest().SetBody(pI).Post(peopleImg + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func UpdatePeopleImage(pI people.People_image) (people.People_image, error) {

	entity := people.People_image{}
	resp, _ := api.Rest().SetBody(pI).Post(peopleImg + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadPeopleImage(id string) (people.People_image, error) {

	entity := people.People_image{}
	resp, _ := api.Rest().Get(peopleImg + "read?id" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func DeletePeopleImage(id string) (people.People_image, error) {

	entity := people.People_image{}
	resp, _ := api.Rest().Get(peopleImg + "delete?id" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadPeopleImages() (people.People_image, error) {

	entity := people.People_image{}
	resp, _ := api.Rest().Get(peopleImg + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
