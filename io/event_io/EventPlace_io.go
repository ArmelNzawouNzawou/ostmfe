package event_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/event"
)

const eventP = api.BASE_URL + "eventP"

func CreateEventPlace(E event.EventPlace) (event.EventPlace, error) {

	entity := event.EventPlace{}
	resp, _ := api.Rest().SetBody(E).Post(eventP + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateEventPlace(E event.EventPlace) (event.EventPlace, error) {

	entity := event.EventPlace{}
	resp, _ := api.Rest().SetBody(E).Post(eventP + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadEventPlace(id string) (event.EventPlace, error) {

	entity := event.EventPlace{}
	resp, _ := api.Rest().Get(eventP + "read?id" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteEventPlace(id string) (event.EventPlace, error) {

	entity := event.EventPlace{}
	resp, _ := api.Rest().Get(eventP + "delete?id" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadEventPlaces() (event.EventPlace, error) {

	entity := event.EventPlace{}
	resp, _ := api.Rest().Get(eventP + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
