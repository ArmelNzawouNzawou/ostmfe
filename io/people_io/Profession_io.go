package people_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/people"
)

const prof = api.BASE_URL + "prof"

func CreateProfession(prf people.Profession) (people.Profession, error) {

	entity := people.Profession{}
	resp, _ := api.Rest().SetBody(prf).Post(prof + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func UpdateProfession(prf people.Profession) (people.Profession, error) {

	entity := people.Profession{}
	resp, _ := api.Rest().SetBody(prf).Post(prof + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadProfession(id string) (people.Profession, error) {

	entity := people.Profession{}
	resp, _ := api.Rest().Get(prof + "read?id" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func DeleteProfession(id string) (people.Profession, error) {

	entity := people.Profession{}
	resp, _ := api.Rest().Get(prof + "delete?id" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadProfessions() (people.Profession, error) {

	entity := people.Profession{}
	resp, _ := api.Rest().Get(prof + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
