package transport_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/transport"
)

const trans = api.BASE_URL + "trans"

func CreateTransport(T transport.TransportType) (transport.TransportType, error) {

	entity := transport.TransportType{}
	resp, _ := api.Rest().SetBody(T).Post(trans + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func UpdateTransport(T transport.TransportType) (transport.TransportType, error) {

	entity := transport.TransportType{}
	resp, _ := api.Rest().SetBody(T).Post(trans + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadTransport(id string) (transport.TransportType, error) {

	entity := transport.TransportType{}
	resp, _ := api.Rest().Get(trans + "read?id" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func DeleteTransport(id string) (transport.TransportType, error) {

	entity := transport.TransportType{}
	resp, _ := api.Rest().Get(trans + "delete?id" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadTransports() (transport.TransportType, error) {

	entity := transport.TransportType{}
	resp, _ := api.Rest().Get(trans + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
