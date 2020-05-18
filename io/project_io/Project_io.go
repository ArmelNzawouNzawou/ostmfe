package project_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/project"
)

const prJ = api.BASE_URL + "prJ"

func CreateProject(P project.Project) (project.Project, error) {

	entity := project.Project{}
	resp, _ := api.Rest().SetBody(P).Post(prJ + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func updateProject(P project.Project) (project.Project, error) {

	entity := project.Project{}
	resp, _ := api.Rest().SetBody(P).Post(prJ + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadProject(id string) (project.Project, error) {

	entity := project.Project{}
	resp, _ := api.Rest().Get(prJ + "read?id" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func DeleteProject(id string) (project.Project, error) {

	entity := project.Project{}
	resp, _ := api.Rest().Get(prJ + "delete?id" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadProjects() (project.Project, error) {

	entity := project.Project{}
	resp, _ := api.Rest().Get(prJ + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
