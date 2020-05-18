package event_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/event"
)

const eventImg = api.BASE_URL + "eventImg"

func CreateEventImg(image event.EventImage) (event.EventImage, error) {

	entity := event.EventImage{}
	resp, _ := api.Rest().SetBody(image).Post(eventImg + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func UpdateEventImg(image event.EventImage) (event.EventImage, error) {

	entity := event.EventImage{}
	resp, _ := api.Rest().SetBody(image).Post(eventImg + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadEventImg(id string) (event.EventImage, error) {

	entity := event.EventImage{}
	resp, _ := api.Rest().Get(eventImg + "read?id" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func DeleteEventImg(id string) (event.EventImage, error) {

	entity := event.EventImage{}
	resp, _ := api.Rest().Get(eventImg + "delete?id" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadEventmgs() (event.EventImage, error) {

	entity := event.EventImage{}
	resp, _ := api.Rest().Get(eventImg + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
