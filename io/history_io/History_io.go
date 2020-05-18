package history_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/history"
)

const story = api.BASE_URL + "story"

func CreateHistory(hist history.History) (history.History, error) {

	entity := history.History{}
	resp, _ := api.Rest().SetBody(hist).Post(story + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func UpdateHistory(hist history.History) (history.History, error) {

	entity := history.History{}
	resp, _ := api.Rest().SetBody(hist).Post(story + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadHistory(id string) (history.History, error) {

	entity := history.History{}
	resp, _ := api.Rest().Get(story + "read?id" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func DeleteHistory(id string) (history.History, error) {

	entity := history.History{}
	resp, _ := api.Rest().Get(story + "delete?id" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadHistorys() (history.History, error) {

	entity := history.History{}
	resp, _ := api.Rest().Get(story + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
