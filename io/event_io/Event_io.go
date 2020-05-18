package event_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/event"
)

const evenT = api.BASE_URL + "evenT"

func CreateEvent(myEvent event.Event) (event.Event, error) {

	entity := event.Event{}
	resp, _ := api.Rest().SetBody(myEvent).Post(evenT + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateEvent(myEvent event.Event) (event.Event, error) {

	entity := event.Event{}
	resp, _ := api.Rest().SetBody(myEvent).Post(evenT + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadEvent(id string) (event.Event, error) {

	entity := event.Event{}
	resp, _ := api.Rest().Get(evenT + "read?id" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteEvent(id string) (event.Event, error) {

	entity := event.Event{}
	resp, _ := api.Rest().Get(evenT + "delete?id" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadEvents() (event.Event, error) {
	entity := event.Event{}
	resp, _ := api.Rest().Get(evenT + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
