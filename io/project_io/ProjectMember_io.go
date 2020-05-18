package project_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/project"
)

const projMemb = api.BASE_URL + "projMemb"

func CreateProjectMember(pm project.ProjectMember) (project.ProjectMember, error) {

	entity := project.ProjectMember{}
	resp, _ := api.Rest().SetBody(pm).Post(projMemb + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateProjectMember(pm project.ProjectMember) (project.ProjectMember, error) {

	entity := project.ProjectMember{}
	resp, _ := api.Rest().SetBody(pm).Post(projMemb + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadProjectMember(id string) (project.ProjectMember, error) {

	entity := project.ProjectMember{}
	resp, _ := api.Rest().Get(projMemb + "read?id" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteProjectMember(id string) (project.ProjectMember, error) {

	entity := project.ProjectMember{}
	resp, _ := api.Rest().Get(projMemb + "delete?id" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadProjectMembers() (project.ProjectMember, error) {

	entity := project.ProjectMember{}
	resp, _ := api.Rest().Get(projMemb + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
