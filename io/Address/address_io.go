package Address

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/address"
)

const addressURL = api.BASE_URL + "address/"

func CreateAddress(myaddress address.Address) (address.Address, error) {
	entity := address.Address{}
	resp, _ := api.Rest().SetBody(myaddress).Post(addressURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateAddress(myaddress address.Address) (address.Address, error) {
	entity := address.Address{}
	resp, _ := api.Rest().SetBody(myaddress).Post(addressURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadAddress(id string) (address.Address, error) {
	entity := address.Address{}
	resp, _ := api.Rest().Get(addressURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteAddress(id string) (address.Address, error) {
	entity := address.Address{}
	resp, _ := api.Rest().Get(addressURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadAddresses() ([]address.Address, error) {
	entity := []address.Address{}
	resp, _ := api.Rest().Get(addressURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
