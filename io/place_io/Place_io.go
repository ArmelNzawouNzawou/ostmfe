package place_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/place"
)

const places = api.BASE_URL + "places"

func CreatePlace(plcs place.Place) (place.Place, error) {

	entity := place.Place{}
	resp, _ := api.Rest().SetBody(plcs).Post(places + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func UpdatePlace(plcs place.Place) (place.Place, error) {

	entity := place.Place{}
	resp, _ := api.Rest().SetBody(plcs).Post(places + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadPlace(id string) (place.Place, error) {

	entity := place.Place{}
	resp, _ := api.Rest().Get(places + "read?id" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func DeletePlace(id string) (place.Place, error) {

	entity := place.Place{}
	resp, _ := api.Rest().Get(places + "delte?id" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadPlaces() (place.Place, error) {

	entity := place.Place{}
	resp, _ := api.Rest().Get(places + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
