package io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain"
)

const mus = api.BASE_URL + "mus"

func CreateMuseum(M museum.Museum) (museum.Museum, error) {

	entity := museum.Museum{}
	resp, _ := api.Rest().SetBody(M).Post(mus + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func UpdateMuseum(M museum.Museum) (museum.Museum, error) {

	entity := museum.Museum{}
	resp, _ := api.Rest().SetBody(M).Post(mus + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadMuseum(id string) (museum.Museum, error) {

	entity := museum.Museum{}
	resp, _ := api.Rest().Get(mus + "read?id" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func DeleteMuseum(id string) (museum.Museum, error) {

	entity := museum.Museum{}
	resp, _ := api.Rest().Get(mus + "delete?id" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadMuseums() (museum.Museum, error) {

	entity := museum.Museum{}
	resp, _ := api.Rest().Get(mus + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
