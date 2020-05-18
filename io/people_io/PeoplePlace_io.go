package people_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/people"
)

const pplPlace = api.BASE_URL + "pplePlace"

func CreatePeoplePlace(pPl people.PeoplePlace) (people.PeoplePlace, error) {
	entity := people.PeoplePlace{}
	resp, _ := api.Rest().SetBody(pPl).Post(pplPlace + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func UpdatePeoplePlace(pPl people.PeoplePlace) (people.PeoplePlace, error) {
	entity := people.PeoplePlace{}
	resp, _ := api.Rest().SetBody(pPl).Post(pplPlace + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadPeoplePlace(id string) (people.PeoplePlace, error) {
	entity := people.PeoplePlace{}
	resp, _ := api.Rest().Get(pplPlace + "read?id" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func DeletePeoplePlace(id string) (people.PeoplePlace, error) {
	entity := people.PeoplePlace{}
	resp, _ := api.Rest().Get(pplPlace + "delete?id" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadPeoplePlaces() (people.PeoplePlace, error) {
	entity := people.PeoplePlace{}
	resp, _ := api.Rest().Get(pplPlace + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
