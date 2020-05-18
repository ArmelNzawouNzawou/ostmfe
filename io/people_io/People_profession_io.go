package people_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/people"
)

const ppleProf = api.BASE_URL + "ppleProf"

func CreatePeopleProfession(pP people.People_profession) (people.People_profession, error) {

	entity := people.People_profession{}
	resp, _ := api.Rest().SetBody(pP).Post(ppleProf + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdatePeopleProfession(pP people.People_profession) (people.People_profession, error) {

	entity := people.People_profession{}
	resp, _ := api.Rest().SetBody(pP).Post(ppleProf + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadPeopleProfession(id string) (people.People_profession, error) {

	entity := people.People_profession{}
	resp, _ := api.Rest().Get(ppleProf + "read?id" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeletePeopleProfession(id string) (people.People_profession, error) {

	entity := people.People_profession{}
	resp, _ := api.Rest().Get(ppleProf + "delete?id" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadPeopleProfessions() (people.People_profession, error) {

	entity := people.People_profession{}
	resp, _ := api.Rest().Get(ppleProf + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
