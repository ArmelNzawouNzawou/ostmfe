package image_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/image"
)

const imG = api.BASE_URL + "imG"

func CreateImage(img image.Images) (image.Images, error) {

	entity := image.Images{}
	resp, _ := api.Rest().SetBody(img).Post(imG + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func updateImage(img image.Images) (image.Images, error) {

	entity := image.Images{}
	resp, _ := api.Rest().SetBody(img).Post(imG + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadImage(id string) (image.Images, error) {

	entity := image.Images{}
	resp, _ := api.Rest().Get(imG + "read?id" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func DeleteImage(id string) (image.Images, error) {

	entity := image.Images{}
	resp, _ := api.Rest().Get(imG + "delete?id" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadImages() (image.Images, error) {

	entity := image.Images{}
	resp, _ := api.Rest().Get(imG + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
