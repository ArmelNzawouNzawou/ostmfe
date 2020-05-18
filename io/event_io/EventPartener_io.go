package event_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/partner"
)

const eventPrtnr = api.BASE_URL + "eventPrtnr"

func CreateEventPartener(prtnr partner.Partner) (partner.Partner, error) {

	entity := partner.Partner{}
	resp, _ := api.Rest().SetBody(prtnr).Post(eventPrtnr + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func UpdateEventPartener(prtnr partner.Partner) (partner.Partner, error) {

	entity := partner.Partner{}
	resp, _ := api.Rest().SetBody(prtnr).Post(eventPrtnr + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadEventPartener(id string) (partner.Partner, error) {

	entity := partner.Partner{}
	resp, _ := api.Rest().Get(eventPrtnr + "read?id" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func DeleteEventPartener(id string) (partner.Partner, error) {

	entity := partner.Partner{}
	resp, _ := api.Rest().Get(eventPrtnr + "delete?id" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadEventParteners() (partner.Partner, error) {

	entity := partner.Partner{}
	resp, _ := api.Rest().Get(eventPrtnr + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
