package project_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/project"
)

const prjPartner = api.BASE_URL + "prjPartner"

func CreateProjectPartner(prjP project.ProjectPartner) (project.ProjectPartner, error) {

	entity := project.ProjectPartner{}
	resp, _ := api.Rest().SetBody(prjP).Post(prjPartner + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func updateProjectPartner(prjP project.ProjectPartner) (project.ProjectPartner, error) {

	entity := project.ProjectPartner{}
	resp, _ := api.Rest().SetBody(prjP).Post(prjPartner + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadProjectPartner(id string) (project.ProjectPartner, error) {

	entity := project.ProjectPartner{}
	resp, _ := api.Rest().Get(prjPartner + "read?id" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func DeleteProjectPartner(id string) (project.ProjectPartner, error) {

	entity := project.ProjectPartner{}
	resp, _ := api.Rest().Get(prjPartner + "delete?id" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadProjectPartners() (project.ProjectPartner, error) {

	entity := project.ProjectPartner{}
	resp, _ := api.Rest().Get(prjPartner + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
