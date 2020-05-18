package collection_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/collection"
)

const collect = api.BASE_URL + "collect"

func CreateCollection(C collection.Collection) (collection.Collection, error) {

	entity := collection.Collection{}
	resp, _ := api.Rest().SetBody(C).Post(collect + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateCollection(C collection.Collection) (collection.Collection, error) {
	entity := collection.Collection{}
	resp, _ := api.Rest().SetBody(C).Post(collect + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadCollection(id string) (collection.Collection, error) {
	entity := collection.Collection{}
	resp, _ := api.Rest().Get(collect + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteCollection(id string) (collection.Collection, error) {
	entity := collection.Collection{}
	resp, _ := api.Rest().Get(collect + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadCollections() (collection.Collection, error) {
	entity := collection.Collection{}
	resp, _ := api.Rest().Get(collect + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
