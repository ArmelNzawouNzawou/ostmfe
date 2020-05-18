package history_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/history"
)

const historyImg = api.BASE_URL + "historyImg"

func CreateHistoryImage(hist history.History_image) (history.History_image, error) {
	entity := history.History_image{}
	resp, _ := api.Rest().SetBody(hist).Post(historyImg + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func UpdateHistoryImage(hist history.History_image) (history.History_image, error) {
	entity := history.History_image{}
	resp, _ := api.Rest().SetBody(hist).Post(historyImg + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadHistoryImage(id string) (history.History_image, error) {
	entity := history.History_image{}
	resp, _ := api.Rest().Get(historyImg + "read?id" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func DeleteHistoryImage(id string) (history.History_image, error) {
	entity := history.History_image{}
	resp, _ := api.Rest().Get(historyImg + "delete?id" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadHistoryImages() (history.History_image, error) {
	entity := history.History_image{}
	resp, _ := api.Rest().Get(historyImg + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
